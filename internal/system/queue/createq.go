package queue

import (
	"bytes"
	"encoding/gob"
	"errors"
	"os"
	"path"
	"strings"
)

// Append a new Queue
func (e *Queue) Append(conf *QueueConf) {
	e.queue = append(e.queue, conf)
}

// Createq creates a new Queue with the provided configuration parameters.
// It checks if the Queue with the same name already exists, returns an error if it does.
// It creates the queue structure on the file system and saves the configuration in a file.
// It adds the new queue to the system and returns nil if the operation is successful.
// If the name parameter is empty, it returns an error.
func (e *Queue) Createq(newq *QueueConf, dir string) error {

	if newq.Name == "" {
		return errors.New("Name can not be null")
	}
	for _, queue := range e.queue {
		if queue.Name == strings.ToLower(newq.Name) {
			return errors.New("Queue already exists")
		}
	}
	//save config
	err := saveInFile(newq, dir)
	if err != nil {
		return err
	}
	//Create dir structure
	err = createStructureQ(path.Join(dir, "/queues/", newq.ID.String()))
	if err != nil {
		return err
	}
	//Add config in system
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.queue = append(e.queue, newq)
	return nil

}

func createStructureQ(basedir string) (err error) {
	err = os.MkdirAll(path.Join(basedir, "/messages"), 0777)
	if err != nil {
		return err
	}
	err = os.MkdirAll(path.Join(basedir, "/dequet"), 0777)
	if err != nil {
		return err
	}
	return nil
}

func saveInFile(q *QueueConf, directory string) (err error) {
	if _, err := os.Stat(path.Join(directory, "/config", "/"+q.ID.String()+".conf")); err != nil {
		os.MkdirAll(path.Join(directory, "/config"), 0777)
	} else {
		return nil
	}
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(q); err != nil {
		return err
	}
	f, err := os.Create(path.Join(directory, "/config/", q.ID.String()+".conf"))
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	return err
}
