package computation

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/flow-go/fvm/handler"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/engine/execution"
	state2 "github.com/onflow/flow-go/engine/execution/state"
	unittest2 "github.com/onflow/flow-go/engine/execution/state/unittest"
	"github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/ledger/complete/wal/fixtures"

	"github.com/onflow/flow-go/engine/execution/computation/committer"
	"github.com/onflow/flow-go/engine/execution/computation/computer"
	"github.com/onflow/flow-go/engine/execution/computation/computer/uploader"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/engine/execution/testutil"
	"github.com/onflow/flow-go/fvm"
	fvmErrors "github.com/onflow/flow-go/fvm/errors"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/mempool/entity"
	"github.com/onflow/flow-go/module/metrics"
	module "github.com/onflow/flow-go/module/mock"
	state_synchronization "github.com/onflow/flow-go/module/state_synchronization/mock"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
)

var scriptLogThreshold = 1 * time.Second

func TestComputeBlockWithStorage(t *testing.T) {
	rt := fvm.NewInterpreterRuntime()

	chain := flow.Mainnet.Chain()

	vm := fvm.NewVirtualMachine(rt)
	execCtx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain))

	privateKeys, err := testutil.GenerateAccountPrivateKeys(2)
	require.NoError(t, err)

	ledger := testutil.RootBootstrappedLedger(vm, execCtx)
	accounts, err := testutil.CreateAccounts(vm, ledger, programs.NewEmptyPrograms(), privateKeys, chain)
	require.NoError(t, err)

	tx1 := testutil.DeployCounterContractTransaction(accounts[0], chain)
	tx1.SetProposalKey(chain.ServiceAddress(), 0, 0).
		SetGasLimit(1000).
		SetPayer(chain.ServiceAddress())

	err = testutil.SignPayload(tx1, accounts[0], privateKeys[0])
	require.NoError(t, err)

	err = testutil.SignEnvelope(tx1, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
	require.NoError(t, err)

	tx2 := testutil.CreateCounterTransaction(accounts[0], accounts[1])
	tx2.SetProposalKey(chain.ServiceAddress(), 0, 0).
		SetGasLimit(1000).
		SetPayer(chain.ServiceAddress())

	err = testutil.SignPayload(tx2, accounts[1], privateKeys[1])
	require.NoError(t, err)

	err = testutil.SignEnvelope(tx2, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
	require.NoError(t, err)

	transactions := []*flow.TransactionBody{tx1, tx2}

	col := flow.Collection{Transactions: transactions}

	guarantee := flow.CollectionGuarantee{
		CollectionID: col.ID(),
		Signature:    nil,
	}

	block := flow.Block{
		Header: &flow.Header{
			View: 42,
		},
		Payload: &flow.Payload{
			Guarantees: []*flow.CollectionGuarantee{&guarantee},
		},
	}

	executableBlock := &entity.ExecutableBlock{
		Block: &block,
		CompleteCollections: map[flow.Identifier]*entity.CompleteCollection{
			guarantee.ID(): {
				Guarantee:    &guarantee,
				Transactions: transactions,
			},
		},
		StartState: unittest.StateCommitmentPointerFixture(),
	}

	me := new(module.Local)
	me.On("NodeID").Return(flow.ZeroID)

	blockComputer, err := computer.NewBlockComputer(vm, execCtx, metrics.NewNoopCollector(), trace.NewNoopTracer(), zerolog.Nop(), committer.NewNoopViewCommitter())
	require.NoError(t, err)

	programsCache, err := NewProgramsCache(10)
	require.NoError(t, err)

	eds := new(state_synchronization.ExecutionDataService)
	eds.On("Add", mock.Anything, mock.Anything).Return(flow.ZeroID, nil, nil)

	eCache := new(state_synchronization.ExecutionDataCIDCache)
	eCache.On("Insert", mock.AnythingOfType("*flow.Header"), mock.AnythingOfType("state_synchronization.BlobTree"))

	engine := &Manager{
		blockComputer: blockComputer,
		me:            me,
		programsCache: programsCache,
		eds:           eds,
		edCache:       eCache,
	}

	view := delta.NewView(ledger.Get)
	blockView := view.NewChild()

	returnedComputationResult, err := engine.ComputeBlock(context.Background(), executableBlock, blockView)
	require.NoError(t, err)

	require.NotEmpty(t, blockView.(*delta.View).Delta())
	require.Len(t, returnedComputationResult.StateSnapshots, 1+1) // 1 coll + 1 system chunk
	assert.NotEmpty(t, returnedComputationResult.StateSnapshots[0].Delta)
	assert.True(t, returnedComputationResult.ComputationUsed > 0)
}

func TestComputeBlock_Uploader(t *testing.T) {

	noopCollector := &metrics.NoopCollector{}

	ledger, err := complete.NewLedger(&fixtures.NoopWAL{}, 10, noopCollector, zerolog.Nop(), complete.DefaultPathFinderVersion)
	require.NoError(t, err)

	me := new(module.Local)
	me.On("NodeID").Return(flow.ZeroID)

	computationResult := unittest2.ComputationResultFixture([][]flow.Identifier{
		{unittest.IdentifierFixture()},
		{unittest.IdentifierFixture()},
	})

	blockComputer := &FakeBlockComputer{
		computationResult: computationResult,
	}

	programsCache, err := NewProgramsCache(10)
	require.NoError(t, err)

	fakeUploader := &FakeUploader{}

	eds := new(state_synchronization.ExecutionDataService)
	eds.On("Add", mock.Anything, mock.Anything).Return(flow.ZeroID, nil, nil)

	eCache := new(state_synchronization.ExecutionDataCIDCache)
	eCache.On("Insert", mock.AnythingOfType("*flow.Header"), mock.AnythingOfType("state_synchronization.BlobTree"))

	manager := &Manager{
		blockComputer: blockComputer,
		me:            me,
		programsCache: programsCache,
		uploaders:     []uploader.Uploader{fakeUploader},
		eds:           eds,
		edCache:       eCache,
	}

	view := delta.NewView(state2.LedgerGetRegister(ledger, flow.StateCommitment(ledger.InitialState())))
	blockView := view.NewChild()

	_, err = manager.ComputeBlock(context.Background(), computationResult.ExecutableBlock, blockView)
	require.NoError(t, err)

	retrievedResult, has := fakeUploader.data[computationResult.ExecutableBlock.ID()]
	require.True(t, has)

	assert.Equal(t, computationResult, retrievedResult)
}

func TestExecuteScript(t *testing.T) {

	logger := zerolog.Nop()

	execCtx := fvm.NewContext(logger)

	me := new(module.Local)
	me.On("NodeID").Return(flow.ZeroID)

	rt := fvm.NewInterpreterRuntime()

	vm := fvm.NewVirtualMachine(rt)

	ledger := testutil.RootBootstrappedLedger(vm, execCtx)

	view := delta.NewView(ledger.Get)

	scriptView := view.NewChild()

	script := []byte(fmt.Sprintf(
		`
			import FungibleToken from %s

			pub fun main() {}
		`,
		fvm.FungibleTokenAddress(execCtx.Chain).HexWithPrefix(),
	))

	eds := new(state_synchronization.ExecutionDataService)
	edCache := new(state_synchronization.ExecutionDataCIDCache)

	engine, err := New(logger,
		metrics.NewNoopCollector(),
		nil,
		me,
		nil,
		vm,
		execCtx,
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		scriptLogThreshold,
		DefaultScriptExecutionTimeLimit,
		nil,
		eds,
		edCache)
	require.NoError(t, err)

	header := unittest.BlockHeaderFixture()
	_, err = engine.ExecuteScript(context.Background(), script, nil, &header, scriptView)
	require.NoError(t, err)
}

func TestExecuteScripPanicsAreHandled(t *testing.T) {

	ctx := fvm.NewContext(zerolog.Nop())

	vm := &PanickingVM{}

	buffer := &bytes.Buffer{}
	log := zerolog.New(buffer)

	header := unittest.BlockHeaderFixture()

	eds := new(state_synchronization.ExecutionDataService)
	edCache := new(state_synchronization.ExecutionDataCIDCache)

	manager, err := New(log,
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		vm,
		ctx,
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		scriptLogThreshold,
		DefaultScriptExecutionTimeLimit,
		nil,
		eds,
		edCache)
	require.NoError(t, err)

	_, err = manager.ExecuteScript(context.Background(), []byte("whatever"), nil, &header, noopView())

	require.Error(t, err)

	require.Contains(t, buffer.String(), "Verunsicherung")
}

func TestExecuteScript_LongScriptsAreLogged(t *testing.T) {

	ctx := fvm.NewContext(zerolog.Nop())

	vm := &LongRunningVM{duration: 2 * time.Millisecond}

	buffer := &bytes.Buffer{}
	log := zerolog.New(buffer)

	header := unittest.BlockHeaderFixture()

	eds := new(state_synchronization.ExecutionDataService)
	edCache := new(state_synchronization.ExecutionDataCIDCache)

	manager, err := New(log,
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		vm,
		ctx,
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		1*time.Millisecond,
		DefaultScriptExecutionTimeLimit,
		nil,
		eds,
		edCache)
	require.NoError(t, err)

	_, err = manager.ExecuteScript(context.Background(), []byte("whatever"), nil, &header, noopView())

	require.NoError(t, err)

	require.Contains(t, buffer.String(), "exceeded threshold")
}

func TestExecuteScript_ShortScriptsAreNotLogged(t *testing.T) {

	ctx := fvm.NewContext(zerolog.Nop())

	vm := &LongRunningVM{duration: 0}

	buffer := &bytes.Buffer{}
	log := zerolog.New(buffer)

	header := unittest.BlockHeaderFixture()

	eds := new(state_synchronization.ExecutionDataService)
	edCache := new(state_synchronization.ExecutionDataCIDCache)

	manager, err := New(log,
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		vm,
		ctx,
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		1*time.Second,
		DefaultScriptExecutionTimeLimit,
		nil,
		eds,
		edCache)
	require.NoError(t, err)

	_, err = manager.ExecuteScript(context.Background(), []byte("whatever"), nil, &header, noopView())

	require.NoError(t, err)

	require.NotContains(t, buffer.String(), "exceeded threshold")
}

type PanickingVM struct{}

func (p *PanickingVM) Run(f fvm.Context, procedure fvm.Procedure, view state.View, p2 *programs.Programs) error {
	panic("panic, but expected with sentinel for test: Verunsicherung ")
}

func (p *PanickingVM) GetAccount(f fvm.Context, address flow.Address, view state.View, p2 *programs.Programs) (*flow.Account, error) {
	panic("not expected")
}

type LongRunningVM struct {
	duration time.Duration
}

func (l *LongRunningVM) Run(f fvm.Context, procedure fvm.Procedure, view state.View, p2 *programs.Programs) error {
	time.Sleep(l.duration)
	// satisfy value marshaller
	if scriptProcedure, is := procedure.(*fvm.ScriptProcedure); is {
		scriptProcedure.Value = cadence.NewVoid()
	}

	return nil
}

func (l *LongRunningVM) GetAccount(f fvm.Context, address flow.Address, view state.View, p2 *programs.Programs) (*flow.Account, error) {
	panic("not expected")
}

type FakeBlockComputer struct {
	computationResult *execution.ComputationResult
}

func (f *FakeBlockComputer) ExecuteBlock(context.Context, *entity.ExecutableBlock, state.View, *programs.Programs) (*execution.ComputationResult, error) {
	return f.computationResult, nil
}

type FakeUploader struct {
	data map[flow.Identifier]*execution.ComputationResult
}

func (f *FakeUploader) Upload(computationResult *execution.ComputationResult) error {
	if f.data == nil {
		f.data = make(map[flow.Identifier]*execution.ComputationResult)
	}
	f.data[computationResult.ExecutableBlock.ID()] = computationResult
	return nil
}

func noopView() *delta.View {
	return delta.NewView(func(_, _, _ string) (flow.RegisterValue, error) {
		return nil, nil
	})
}

func TestExecuteScriptTimeout(t *testing.T) {

	timeout := 1 * time.Millisecond
	manager, err := New(
		zerolog.Nop(),
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		fvm.NewVirtualMachine(fvm.NewInterpreterRuntime()),
		fvm.NewContext(zerolog.Nop()),
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		DefaultScriptLogThreshold,
		timeout,
		nil,
		nil,
		nil)

	require.NoError(t, err)

	script := []byte(`
	pub fun main(): Int {
		var i = 0
		while i < 10000 {
			i = i + 1
		}
		return i
	}
	`)

	header := unittest.BlockHeaderFixture()
	value, err := manager.ExecuteScript(context.Background(), script, nil, &header, noopView())

	require.Error(t, err)
	require.Nil(t, value)
	require.Contains(t, err.Error(), fvmErrors.ErrCodeScriptExecutionTimedOutError.String())
}

func TestExecuteScriptCancelled(t *testing.T) {

	timeout := 30 * time.Second
	manager, err := New(
		zerolog.Nop(),
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		fvm.NewVirtualMachine(fvm.NewInterpreterRuntime()),
		fvm.NewContext(zerolog.Nop()),
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		DefaultScriptLogThreshold,
		timeout,
		nil,
		nil,
		nil)

	require.NoError(t, err)

	script := []byte(`
	pub fun main(): Int {
		var i = 0
		var j = 0 
		while i < 10000000 {
			i = i + 1
			j = i + j
		}
		return i
	}
	`)

	var value []byte
	var wg sync.WaitGroup
	reqCtx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		header := unittest.BlockHeaderFixture()
		value, err = manager.ExecuteScript(reqCtx, script, nil, &header, noopView())
		wg.Done()
	}()
	cancel()
	wg.Wait()
	require.Nil(t, value)
	require.Contains(t, err.Error(), fvmErrors.ErrCodeScriptExecutionCancelledError.String())
}

func TestScriptStorageMutationsDiscarded(t *testing.T) {

	timeout := 1 * time.Millisecond
	vm := fvm.NewVirtualMachine(fvm.NewInterpreterRuntime())
	chain := flow.Mainnet.Chain()
	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain))
	manager, _ := New(
		zerolog.Nop(),
		metrics.NewNoopCollector(),
		nil,
		nil,
		nil,
		vm,
		ctx,
		DefaultProgramsCacheSize,
		committer.NewNoopViewCommitter(),
		DefaultScriptLogThreshold,
		timeout,
		nil,
		nil,
		nil)
	view := testutil.RootBootstrappedLedger(vm, ctx)
	programs := programs.NewEmptyPrograms()
	st := state.NewState(view)
	sth := state.NewStateHolder(st)
	env := fvm.NewScriptEnvironment(context.Background(), ctx, vm, sth, programs)

	// Create an account private key.
	privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
	require.NoError(t, err)

	// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
	accounts, err := testutil.CreateAccounts(vm, view, programs, privateKeys, chain)
	require.NoError(t, err)
	account := accounts[0]
	address := cadence.NewAddress(account)
	commonAddress, _ := common.HexToAddress(address.Hex())

	script := []byte(`
	pub fun main(account: Address) {
		let acc = getAuthAccount(account)
		acc.save(3, to: /storage/x)
	}
	`)

	header := unittest.BlockHeaderFixture()
	scriptView := view.NewChild()
	_, err = manager.ExecuteScript(context.Background(), script, [][]byte{jsoncdc.MustEncode(address)}, &header, scriptView)

	require.NoError(t, err)

	v, err := vm.Runtime.ReadStored(
		commonAddress,
		cadence.NewPath("storage", "x"),
		runtime.Context{Interface: env},
	)

	// the save should not update account storage by writing the delta from the child view back to the parent
	require.NoError(t, err)
	require.Equal(t, nil, v)
}

func Test_EventEncodingFailsOnlyTxAndCarriesOn(t *testing.T) {
	rt := fvm.NewInterpreterRuntime()
	chain := flow.Mainnet.Chain()
	vm := fvm.NewVirtualMachine(rt)

	eventEncoder := &testingEventEncoder{
		CadenceEventEncoder: *handler.NewCadenceEventEncoder(),
	}

	execCtx := fvm.NewContext(zerolog.Nop(),
		fvm.WithChain(chain),
		fvm.WithTransactionProcessors(
			fvm.NewTransactionInvoker(zerolog.Nop(),
				fvm.WithFlowEventHandleOptions(handler.WithEncoder(eventEncoder)))))

	privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
	require.NoError(t, err)
	ledger := testutil.RootBootstrappedLedger(vm, execCtx)
	accounts, err := testutil.CreateAccounts(vm, ledger, programs.NewEmptyPrograms(), privateKeys, chain)
	require.NoError(t, err)

	// setup transactions
	account := accounts[0]
	privKey := privateKeys[0]
	// tx1 deploys contract version 1
	tx1 := testutil.DeployEventContractTransaction(account, chain, 1)
	prepareTx(t, tx1, account, privKey, 0, chain)

	// tx2 emits event which will fail encoding
	tx2 := testutil.CreateEmitEventTransaction(account, account)
	prepareTx(t, tx2, account, privKey, 1, chain)

	// tx3 emits event that will work fine
	tx3 := testutil.CreateEmitEventTransaction(account, account)
	prepareTx(t, tx3, account, privKey, 2, chain)

	transactions := []*flow.TransactionBody{tx1, tx2, tx3}

	col := flow.Collection{Transactions: transactions}

	guarantee := flow.CollectionGuarantee{
		CollectionID: col.ID(),
		Signature:    nil,
	}

	block := flow.Block{
		Header: &flow.Header{
			View: 26,
		},
		Payload: &flow.Payload{
			Guarantees: []*flow.CollectionGuarantee{&guarantee},
		},
	}

	executableBlock := &entity.ExecutableBlock{
		Block: &block,
		CompleteCollections: map[flow.Identifier]*entity.CompleteCollection{
			guarantee.ID(): {
				Guarantee:    &guarantee,
				Transactions: transactions,
			},
		},
		StartState: unittest.StateCommitmentPointerFixture(),
	}

	me := new(module.Local)
	me.On("NodeID").Return(flow.ZeroID)

	blockComputer, err := computer.NewBlockComputer(vm, execCtx, metrics.NewNoopCollector(), trace.NewNoopTracer(), zerolog.Nop(), committer.NewNoopViewCommitter())
	require.NoError(t, err)

	programsCache, err := NewProgramsCache(10)
	require.NoError(t, err)

	eds := new(state_synchronization.ExecutionDataService)
	eds.On("Add", mock.Anything, mock.Anything).Return(flow.ZeroID, nil, nil)

	edCache := new(state_synchronization.ExecutionDataCIDCache)
	edCache.On("Insert", mock.AnythingOfType("*flow.Header"), mock.AnythingOfType("BlobTree"))

	engine := &Manager{
		blockComputer: blockComputer,
		me:            me,
		programsCache: programsCache,
		eds:           eds,
		edCache:       edCache,
	}

	view := delta.NewView(ledger.Get)
	blockView := view.NewChild()

	returnedComputationResult, err := engine.ComputeBlock(context.Background(), executableBlock, blockView)
	require.NoError(t, err)

	require.Len(t, returnedComputationResult.Events, 2)             // 1 collection + 1 system chunk
	require.Len(t, returnedComputationResult.TransactionResults, 4) // 2 txs + 1 system tx

	require.Empty(t, returnedComputationResult.TransactionResults[0].ErrorMessage)
	require.Contains(t, returnedComputationResult.TransactionResults[1].ErrorMessage, "I failed encoding")
	require.Empty(t, returnedComputationResult.TransactionResults[2].ErrorMessage)

	// first event should be contract deployed
	assert.EqualValues(t, "flow.AccountContractAdded", returnedComputationResult.Events[0][0].Type)

	// second event should come from tx3 (index 2)  as tx2 (index 1) should fail encoding
	hasValidEventValue(t, returnedComputationResult.Events[0][1], 1)
	assert.Equal(t, returnedComputationResult.Events[0][1].TransactionIndex, uint32(2))
}

type testingEventEncoder struct {
	handler.CadenceEventEncoder
	calls int
}

func (e *testingEventEncoder) Encode(event cadence.Event) ([]byte, error) {
	defer func() {
		e.calls++
	}()

	fmt.Printf("%s\n", event)

	if e.calls == 1 {
		return nil, fmt.Errorf("I failed encoding")
	}
	return e.CadenceEventEncoder.Encode(event)

}
