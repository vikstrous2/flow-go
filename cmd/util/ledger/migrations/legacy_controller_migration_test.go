package migrations

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go-sdk"

	"github.com/onflow/flow-go/engine/execution/state"
	fvmstate "github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/utils"
)

const LegacyKeyPartController = 1

func createPayloadKeyWithLegacyController(a flow.Address, key string, emptyController bool) ledger.Key {
	if emptyController {
		return ledger.Key{
			KeyParts: []ledger.KeyPart{
				ledger.NewKeyPart(state.KeyPartOwner, a.Bytes()),
				ledger.NewKeyPart(LegacyKeyPartController, []byte("")),
				ledger.NewKeyPart(state.KeyPartKey, []byte(key)),
			},
		}
	}
	return ledger.Key{
		KeyParts: []ledger.KeyPart{
			ledger.NewKeyPart(state.KeyPartOwner, a.Bytes()),
			ledger.NewKeyPart(LegacyKeyPartController, a.Bytes()),
			ledger.NewKeyPart(state.KeyPartKey, []byte(key)),
		},
	}
}

func createMigratedPayloadKey(a flow.Address, key string) ledger.Key {
	return ledger.Key{
		KeyParts: []ledger.KeyPart{
			ledger.NewKeyPart(state.KeyPartOwner, a.Bytes()),
			ledger.NewKeyPart(state.KeyPartKey, []byte(key)),
		},
	}
}

func TestLegacyControllerMigration(t *testing.T) {
	mig := LegacyControllerMigration{
		Logger: zerolog.Logger{},
	}

	address1 := flow.HexToAddress("0x1")
	address2 := flow.HexToAddress("0x2")

	payloads := []ledger.Payload{
		{Key: createPayloadKeyWithLegacyController(address1, fvmstate.KeyStorageUsed, false), Value: utils.Uint64ToBinary(1)},
		{Key: createPayloadKeyWithLegacyController(address1, fvmstate.ContractKey("CoreContract"), true), Value: utils.Uint64ToBinary(2)},
		{Key: createPayloadKeyWithLegacyController(address1, fvmstate.KeyContractNames, true), Value: utils.Uint64ToBinary(3)},
		{Key: createPayloadKeyWithLegacyController(address2, fvmstate.KeyPublicKey(1), true), Value: utils.Uint64ToBinary(4)},
		{Key: createPayloadKeyWithLegacyController(address2, fvmstate.KeyPublicKeyCount, true), Value: utils.Uint64ToBinary(4)},
	}

	expectedKeys := []ledger.Key{
		createMigratedPayloadKey(address1, fvmstate.KeyStorageUsed),
		createMigratedPayloadKey(address1, fvmstate.ContractKey("CoreContract")),
		createMigratedPayloadKey(address1, fvmstate.KeyContractNames),
		createMigratedPayloadKey(address2, fvmstate.KeyPublicKey(1)),
		createMigratedPayloadKey(address2, fvmstate.KeyPublicKeyCount),
	}

	newPayloads, err := mig.Migrate(payloads)
	require.NoError(t, err)
	require.Equal(t, len(payloads), len(newPayloads))

	for i, p := range newPayloads {
		require.Equal(t, expectedKeys[i], p.Key)
	}

}
