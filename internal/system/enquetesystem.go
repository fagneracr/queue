package system

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"os"
	"path"
)

// Enq - Init global variable
var Enq *System

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
func InitQ(config Conf) (*System, error) {
	Enq = new(System)
	Enq.config = config

	//Read queues config
	files, err := OSReadDir(path.Join(config.Directory, "/config"))
	if err != nil {
		// New queue system, nor exists
		err = os.Mkdir(path.Join(config.Directory, "/config"), os.ModePerm)
		Enq.config = config
		return Enq, err
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
		// Append a queue config to system.
		Enq.Queue = append(Enq.Queue, &q)
	}
	return Enq, errorout

}

// /*Save - Create a new queue in sysstem*/
// func Save(in Identify) error {

// 	if Enq.mutex.TryLock() {
// 		defer Enq.mutex.Unlock()
// 		for _, i := range Enq.ident {
// 			if i.Name == in.Name {
// 				return errors.New("Configuration name alred exists")
// 			}
// 		}
// 		Enq.ident = append(Enq.ident, in)
// 		saveFileGobAndGz(in)

// 	} else {
// 		return errors.New("Try again later, server busy")
// 	}

// 	return nil
// 	// Message add the id for file queue, after broken get message in order.
// 	// Save file and return uuid or error
// }

func (e *System) Reload() {
	// zero e
	// init again
}

func (e *System) Enquete(msg interface{}, header interface{}, name string) {

	// for _, i := range e.ident {
	// 	if i.Name == name {
	// 		i.Messages = append(i.Messages, message{Msg: msg, Header: header, ID: i.NextID})
	// 		i.NextID++
	// 	}

	// }

}

func OSReadDir(root string) ([]string, error) {
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

func decodeFile(todecode io.Reader) (queueConf, error) {
	var n2 queueConf
	if err := gob.NewDecoder(todecode).Decode(&n2); err != nil {
		return n2, err
	}
	return n2, nil
}
