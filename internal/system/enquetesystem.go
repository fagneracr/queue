package system

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Enq- Init global variable
var Enq *System

/*InitQ - Init a queue system*/
func InitQ() (*System, error) {
	Enq = new(System)
	cfgs, err := OSReadDir("./config")
	if err != nil {
		os.Mkdir("./config", os.ModePerm)
		return nil, err
	}
	var errorout error
	// find config already created queue
	for _, cf := range cfgs {
		out, err := os.Open(path.Join("./config", cf))
		if err != nil {
			errorout = errors.Join(errorout, errors.New("File is broken: "+cf+" with error: "+err.Error()))
			continue
		}
		r, err := gzip.NewReader(out)
		if err != nil {
			errorout = errors.Join(errorout, errors.New("This file can not be decompressed: "+err.Error()))
			continue

		}
		s, _ := ioutil.ReadAll(r)
		reader := bytes.NewReader(s)
		config, err := decodeFile(reader)
		if err != nil {
			errorout = errors.Join(errorout, errors.New("Can not decode the file: "+err.Error()))
			continue
		}
		_ = config
		//Enq.config = append(Enq.ident, config)
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

func decodeFile(todecode io.Reader) (Identify, error) {
	var n2 Identify
	if err := gob.NewDecoder(todecode).Decode(&n2); err != nil {
		return n2, err
	}
	return n2, nil
}
