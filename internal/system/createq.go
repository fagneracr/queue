package system

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

/*Createq - Create a Queue in the system, there are some parameter you con send in config.*/
func (e *System) Createq(config ConfigQueue) error {
	if config.Name == "" {
		return errors.New("Name can not be null")
	}
	var q queueConf
	q.Name = strings.ToLower(config.Name)
	q.maxSize = config.MaxSize
	q.ID = uuid.New()
	q.persistent = config.Persistent
	q.NextID = 1
	q.createDate = time.Now()
	q.TTL = time.Duration(time.Duration(config.TTL) * time.Minute)
	q.mutex = sync.Mutex{}
	for _, v := range config.Variable {
		q.Variable = append(q.Variable, v)
	}
	err := saveInFile(q)
	if err != nil {
		return err
	}
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.Queue = append(e.Queue, &q)
	//save in file persistent

	return nil

}

func saveInFile(q queueConf) (err error) {
	if _, err := os.Stat(path.Join("./config/queues", q.ID.String())); err != nil {
		os.MkdirAll("./config/queues", 0777)
	} else {
		return nil
	}
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(q); err != nil {
		return err
	}
	f, err := os.Create(path.Join("./config/queues", q.ID.String()+".gz"))
	if err != nil {
		return err
	}
	defer f.Close()
	w := gzip.NewWriter(f)
	w.Write(buf.Bytes())
	w.Close()
	return err

}
