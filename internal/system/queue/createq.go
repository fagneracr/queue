package queue

import (
	"bytes"
	"encoding/gob"
	"errors"
	"os"
	"path"
	"strings"
)

func (e *Queue) Append(conf *QueueConf) {
	e.queue = append(e.queue, conf)
}

// Createq creates a new Queue with the provided configuration parameters.
// It checks if the Queue with the same name already exists, returns an error if it does.
// It creates the queue structure on the file system and saves the configuration in a file.
// It adds the new queue to the system and returns nil if the operation is successful.
// If the name parameter is empty, it returns an error.
func (e *Queue) Createq(qnew interface{}, dir string) error {
	if _, ok := qnew.(Queue); !ok {
		return errors.New("does not working conversion")
	}
	newq := qnew.(QueueConf)
	if newq.Name == "" {
		return errors.New("Name can not be null")
	}
	for _, queue := range e.queue {
		if queue.Name == strings.ToLower(newq.Name) {
			return errors.New("Queue already exists")
		}
	}
	// var q queueConf
	// q.Name = strings.ToLower(newq.Name)
	// q.maxSize = newq.MaxSize
	// q.ID = uuid.New()
	// q.persistent = newq.Persistent
	// q.NextID = 1
	// q.createDate = time.Now()
	// q.TTL = time.Duration(time.Duration(newq.TTL) * time.Minute)
	// q.mutex = sync.Mutex{}
	// for _, v := range newq.Variable {
	// 	var variable Variable
	// 	variable.Key = v.Key
	// 	variable.Value = v.Value
	// 	q.Variable = append(q.Variable, variable)
	// }
	//save config
	err := saveInFile(&newq, dir)
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
	e.queue = append(e.queue, &newq)
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
