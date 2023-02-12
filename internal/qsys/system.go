package qsys

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"go-queue/internal/qsys/queue"
	"io"
	"log"
	"os"
	"path"
	"time"
)

// Enq - Init global variable
var Enq *QSys

// Conf - Parameter to Initialize system
type Conf struct {
	Directory string
	//max size in direcoty.
	//log directory

}

/*
InitQ initializes a new System and reads the configuration files for the queues.
If the config directory does not exist, it is created. If there are any errors
in reading or decoding the configuration files, they are returned as a single error.

Returns:

	*System: A pointer to the new System object that was created.
	error: An error object that may contain any errors that occurred while reading
	       or decoding the configuration files.

Example:

	sys, err := InitQ()
	if err != nil {
	    fmt.Printf("Error initializing System: %v\n", err)
	    return
	}

Use the sys object to perform queue operations.
*/
func InitQ(config Conf) error {
	Enq = new(QSys)
	Enq.config = config
	Enq.queue = new(queue.Queue)

	//Read queues config
	files, err := oSReadDir(path.Join(config.Directory, "/config"))
	if err != nil {
		// New queue system, nor exists
		err = os.Mkdir(path.Join(config.Directory, "/config"), os.ModePerm)
		Enq.config = config
		return err
	}

	var errorout error
	configdir := path.Join(path.Join(config.Directory, "/config"))
	// find config already created queue
	for _, file := range files {
		out, err := os.Open(path.Join(configdir, file))
		if err != nil {
			errorout = errors.Join(errorout, errors.New("File is broken: "+file+" with error: "+err.Error()))
			continue
		}
		filecontent, err := io.ReadAll(out)
		reader := bytes.NewReader(filecontent)
		q, err := decodeFile(reader)
		if err != nil {
			errorout = errors.Join(errorout, errors.New("Can not decode the file: "+err.Error()))
			continue
		}
		Enq.queue.Append(&q)
	}
	return errorout

}

// CreateQ creates a new queue with the provided configuration in the System.
//
// Parameters:
// - in: ConfigQueue - configuration of the queue to be created.
//
// Returns:
// - error - an error if the creation of the queue fails.
//
// Example:
//
//	sys := &System{}
//	config := ConfigQueue{Name: "myQueue", TTL: 30, Persistent: true, Variable: []Variable{{Key: "key1", Value: "value1"}, {Key: "key2", Value: "value2"}}}
//	err := sys.CreateQ(config)
//	if err != nil {
//	  fmt.Println("Error creating queue: ", err)
//	}
func CreateQ(in ConfigQueue) (err error) {
	var toCreaate queue.QueueConf
	toCreaate.Name = in.Name
	//corrigir
	toCreaate.TTL = time.Duration(in.TTL)
	toCreaate.Persistent = in.Persistent
	for _, i := range in.Variable {
		var v queue.Variable
		v.Key = i.Key
		v.Value = i.Key
		toCreaate.Variable = append(toCreaate.Variable, v)
	}
	err = Enq.queue.Createq(&toCreaate, Enq.config.Directory)
	return err

}

// DeleteQ deletes a queue with the given `name` from the `System`.
// Returns an error if the queue does not exist or if there was an error deleting
// the queue's struct and config files.
func DeleteQ(name string) (err error) {
	err = Enq.queue.DeleteQ(name, Enq.config.Directory)
	return err
}

// ExistsQ checks if a queue with the given `name` exists in the `System`.
// Returns true if the queue exists, false otherwise.
func ExistsQ(name string) bool {
	return Enq.queue.Exists(name)
}

// ListQ returns a slice of `ConfigQueue` structs representing the configuration
// of the named queues or all queues in the `System`. If one or more queue names
// are provided as arguments, only those queues will be included in the output.
// Otherwise, all queues in the `System` will be included.
func ListQ(names ...string) (confq []ConfigQueue) {
	s := Enq
	i := s.queue.ListQ(names...)
	for _, v := range i {
		var variable []Variable
		for _, k := range v.Variable {
			variable = append(variable, Variable{
				Key:   k.Key,
				Value: k.Value,
			})
		}
		confq = append(confq, ConfigQueue{
			Name:       v.Name,
			Persistent: v.Persistent,
			TTL:        int(v.TTL),
			MaxSize:    0,
			Variable:   variable,
		})
	}
	return confq

}

func oSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func saveFileGobAndGz(input Identify) {
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(input); err != nil {
		panic(err)
	}
	f, err := os.Create(path.Join("./config", "config"+input.Name+".gz"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := gzip.NewWriter(f)
	w.Write(buf.Bytes())
	w.Close()

}

func decodeFile(todecode io.Reader) (queue.QueueConf, error) {
	var n2 queue.QueueConf
	if err := gob.NewDecoder(todecode).Decode(&n2); err != nil {
		return n2, err
	}
	return n2, nil
}
