package system

import (
	"errors"
	"os"
	"path"
	"strings"
)

var pos int

// DeleteQ deletes a queue with the given `name` from the `System`.
// Returns an error if the queue does not exist or if there was an error deleting
// the queue's struct and config files.
func (e *System) DeleteQ(name string) (err error) {
	// Check if the queue with the given `name` exists in the `System`.
	if e.Exists(name) {
		// If the queue exists, obtain a lock on the `System`'s mutex to ensure
		// exclusive access to the `Queue` slice.
		e.mutex.Lock()
		defer e.mutex.Unlock()
		// Iterate over each `Queue` in the `System`.
		for pos, q := range e.Queue {
			if q.Name == strings.ToLower(name) {
				// If the current `Queue`'s name matches the provided `name`, delete
				// the struct and config files for the `Queue` and remove the `Queue`
				// from the `System`'s `Queue` slice.
				err = deleteStructureQ(e.config.Directory+"/", q.ID.String())
				if err != nil {
					return err
				}
				e.Queue = append(e.Queue[:pos], e.Queue[pos+1:]...)
				if pos > 0 {
					pos = pos - 1
				}
				continue
			}
			pos++
		}
	} else {
		// If the queue does not exist, return an error.
		err = errors.New("Queue not found")
	}
	return err
}

func deleteStructureQ(basedir string, id string) (err error) {
	e := os.Remove(path.Join(basedir, "config/", id+".conf"))
	errors.Join(e, os.RemoveAll(path.Join(basedir, "/queues/", id)))
	return e
}
