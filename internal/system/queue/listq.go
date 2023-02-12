package queue

import (
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Queue struct {
	queue []*QueueConf
	mutex sync.Mutex
}

type message struct {
	Msg    interface{}
	Header interface{}
	ID     int64
}

// Variable component config
type Variable struct {
	Key   string
	Value string
}

/*queue internal configuration*/
type QueueConf struct {
	Name       string
	ID         uuid.UUID
	Persistent bool
	TTL        time.Duration
	maxSize    int64
	createDate time.Time
	NextID     int64
	Variable   []Variable
	Messages   []*message
	mutex      sync.Mutex
}

// ListQ returns a slice of `ConfigQueue` structs representing the configuration
// of the named queues or all queues in the `System`. If one or more queue names
// are provided as arguments, only those queues will be included in the output.
// Otherwise, all queues in the `System` will be included.
func (e *Queue) ListQ(name ...string) (confq []*QueueConf) {
	if len(name) > 0 {
		// If queue names are provided, iterate over each name and append
		// the corresponding `ConfigQueue` to the output slice.
		for _, n := range name {
			c := e.findq(n)
			if reflect.DeepEqual(QueueConf{}, &c) {
				continue
			}
			confq = append(confq, c)
		}
	} else {
		// If no queue names are provided, append the configuration of all
		// queues in the `System` to the output slice.
		for _, q := range e.queue {
			confq = append(confq, q)
		}
	}
	// Return the output slice of `ConfigQueue` structs.
	return confq
}

// Exists checks if a queue with the given `name` exists in the `System`.
// Returns true if the queue exists, false otherwise.
func (e *Queue) Exists(name string) bool {
	// Iterate over each `Queue` in the `System`.
	for _, q := range e.queue {
		// If the `name` of the current `Queue` matches the provided `name`,
		// the queue exists, so return true.
		if q.Name == name {
			return true
		}
	}
	// If no matching `Queue` is found, the queue does not exist, so return false.
	return false
}

// func produceOutConfig(e *queueConf) ConfigQueue {
// 	return interface{}{
// 		"MaxSize":    e.maxSize,
// 		"Name":       e.Name,
// 		"Persistent": e.persistent,
// 		"TTL":        int(e.TTL),
// 		"Variable":   e.Variable,
// 	}
// }

func (e *Queue) findq(name string) (confq *QueueConf) {
	for _, q := range e.queue {
		if strings.ToLower(name) == q.Name {
			return q
		}
	}
	return confq

}
