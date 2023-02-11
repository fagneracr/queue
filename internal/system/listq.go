package system

import (
	"reflect"
	"strings"
)

// ListQ returns a slice of `ConfigQueue` structs representing the configuration
// of the named queues or all queues in the `System`. If one or more queue names
// are provided as arguments, only those queues will be included in the output.
// Otherwise, all queues in the `System` will be included.
func (e *System) ListQ(name ...string) (confq []ConfigQueue) {
	if len(name) > 0 {
		// If queue names are provided, iterate over each name and append
		// the corresponding `ConfigQueue` to the output slice.
		for _, n := range name {
			c := e.findq(n)
			if reflect.DeepEqual(queueConf{}, &c) {
				continue
			}
			confq = append(confq, produceOutConfig(c))
		}
	} else {
		// If no queue names are provided, append the configuration of all
		// queues in the `System` to the output slice.
		for _, q := range e.Queue {
			confq = append(confq, produceOutConfig(q))
		}
	}
	// Return the output slice of `ConfigQueue` structs.
	return confq
}

// Exists checks if a queue with the given `name` exists in the `System`.
// Returns true if the queue exists, false otherwise.
func (e *System) Exists(name string) bool {
	// Iterate over each `Queue` in the `System`.
	for _, q := range e.Queue {
		// If the `name` of the current `Queue` matches the provided `name`,
		// the queue exists, so return true.
		if q.Name == name {
			return true
		}
	}
	// If no matching `Queue` is found, the queue does not exist, so return false.
	return false
}

func produceOutConfig(e *queueConf) ConfigQueue {
	return ConfigQueue{
		MaxSize:    e.maxSize,
		Name:       e.Name,
		Persistent: e.persistent,
		TTL:        int(e.TTL),
		Variable:   e.Variable,
	}
}

func (e *System) findq(name string) (confq *queueConf) {
	for _, q := range e.Queue {
		if strings.ToLower(name) == q.Name {
			return q
		}
	}
	return confq

}
