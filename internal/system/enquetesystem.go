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
	"sync"
)

type Enqueted struct {
	ident []Identify
	mutex sync.Mutex
}

type Identify struct {
	ID         string  `json:"id,omitempty"`
	Name       string  `json:"name"`
	FileConfig string  `json:"fileconfig,omitempty"`
	MsgType    Msgtype `json:"msgtype"`
	Messages   []message
	NextID     int64
}
type Msgtype struct {
	Mail mail `json:"mail,omitempty"`
	Mft  mft  `json:"mft,omitempty"`
}

type mail struct {
	To          string `json:"to,,omitempty"`
	Cc          string `json:"cc,,omitempty"`
	Cco         string `json:"cco,,omitempty"`
	MessageBody string `json:"messagebody,,omitempty"`
}
type mft struct {
	Host     string
	Port     int
	User     string
	Passwd   string // bcript
	SenDir   string
	Protocol string
}
type message struct {
	Msg    interface{}
	Header interface{}
	ID     int64
}

var Enq *Enqueted

func InitQ() {
	// Change to parssing config on init.
	Enq = new(Enqueted)
	cfgs, err := OSReadDir("./config")
	if err != nil {
		return
	}
	for _, cf := range cfgs {
		out, _ := os.Open(path.Join("./config", cf))
		r, err := gzip.NewReader(out)
		if err != nil {

		}
		s, _ := ioutil.ReadAll(r)

		reader := bytes.NewReader(s)
		config := decodeFile(reader)
		Enq.ident = append(Enq.ident, config)
	}

}

func (e *Enqueted) Save(input interface{}) error {
	if _, err := input.(Identify); !err {
		return errors.New("Can not convert received input")

	}
	in := input.(Identify)
	if e.mutex.TryLock() {
		defer e.mutex.Unlock()
		for _, i := range e.ident {
			if i.Name == in.Name {
				return errors.New("Configuration name alred exists")
			}
		}
		e.ident = append(e.ident, in)
		saveFileGobAndGz(in)

	} else {
		return errors.New("Try again later, server busy")
	}

	return nil
	// Message add the id for file queue, after broken get message in order.
	// Save file and return uuid or error
}

func (e *Enqueted) Reload() {
	// zero e
	// init again
}

func (e *Enqueted) Enquete(msg interface{}, header interface{}, name string) {

	for _, i := range e.ident {
		if i.Name == name {
			i.Messages = append(i.Messages, message{Msg: msg, Header: header, ID: i.NextID})
			i.NextID++
		}

	}

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

func decodeFile(todecode io.Reader) Identify {
	var n2 Identify
	if err := gob.NewDecoder(todecode).Decode(&n2); err != nil {
		panic(n2)
	}
	return n2
}
