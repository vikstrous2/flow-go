package queue_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dapperlabs/flow-go/network/gossip/libp2p/queue"
)

// TestRetrievalByPriority tests that message can be retrieved in priority order
func TestRetrievalByPriority(t *testing.T) {
	// create a map of messages -> priority with messages assigned random priorities
	messages := createMessages(1000, randomPriority)
	testQueue(t, messages)
}

// TestRetrievalByInsertionOrder tests that messages with the same priority can be retrieved in insertion order
func TestRetrievalByInsertionOrder(t *testing.T) {

	// create a map of messages -> priority with messages assigned fixed priorities
	messages := createMessages(1000, fixedPriority)
	testQueue(t, messages)
}

func testQueue(t *testing.T, messages map[string]queue.Priority) {

	// create the priority function
	var priorityFunc queue.MessagePriorityFunc = func(message interface{}) queue.Priority {
		return messages[message.(string)]
	}

	// create queues for each priority to check expectations later
	queues := make(map[queue.Priority][]string)
	for p := queue.Low_Priority; p <= queue.High_Priority; p++ {
		queues[p] = make([]string, 0)
	}

	// create the queue
	mq := queue.NewMessageQueue(priorityFunc)

	// insert all elements in the queue
	for msg, p := range messages {

		err := mq.Insert(msg)
		assert.NoError(t, err)

		// remember insertion order to check later
		queues[p] = append(queues[p], msg)

		time.Sleep(1 * time.Millisecond)
	}

	// create a slice of the expected messages in the order in which they are expected
	var expectedMessages []string
	for p := queue.High_Priority; p >= queue.Low_Priority; p-- {
		expectedMessages = append(expectedMessages, queues[p]...)
	}

	// check queue length
	assert.Equal(t, len(expectedMessages), mq.Len())

	// check that elements are retrieved in order
	for i := 0; i < len(expectedMessages); i++ {

		item := mq.Remove()

		assert.Equal(t, expectedMessages[i], item.(string))
	}
}

func BenchmarkPush(b *testing.B) {
	b.StopTimer()
	var mq = queue.NewMessageQueue(randomPriority)
	for i := 0; i < b.N; i++ {
		err := mq.Insert("test")
		if err != nil {
			b.Error(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := mq.Insert("test")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkPop(b *testing.B) {
	b.StopTimer()
	var mq = queue.NewMessageQueue(randomPriority)
	for i := 0; i < b.N; i++ {
		err := mq.Insert("test")
		if err != nil {
			b.Error(err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mq.Remove()
	}
}

func createMessages(messageCnt int, priorityFunc queue.MessagePriorityFunc) map[string]queue.Priority {
	msgPrefix := "message"
	// create a map of messages -> priority
	messages := make(map[string]queue.Priority, messageCnt)

	for i := 0; i < messageCnt; i++ {
		// choose a random priority
		p := priorityFunc(nil)
		// create a message
		msg := msgPrefix + strconv.Itoa(i)
		messages[msg] = p
	}

	return messages
}

func randomPriority(_ interface{}) queue.Priority {
	rand.Seed(time.Now().UnixNano())
	p := rand.Intn(int(queue.High_Priority-queue.Low_Priority+1)) + int(queue.Low_Priority)
	return queue.Priority(p)
}

func fixedPriority(_ interface{}) queue.Priority {
	return queue.Priority_5
}
