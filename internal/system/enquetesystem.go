package system

import (
	"encoding/json"
	"os"
	"path"

	"github.com/google/uuid"
)

type Enqueted struct {
	ident []identify
}

type identify struct {
	ID         uuid.UUID
	Name       string
	FileConfig string
	message    []message
	nextID     int64
}
type message struct {
	Msg    interface{}
	Header interface{}
	ID     int64
}

var Enq *Enqueted

func InitQ() {
	Enq = new(Enqueted)
	cfgs, err := OSReadDir("./config")
	if err != nil {
		return
	}
	for _, cf := range cfgs {
		out, _ := os.ReadFile(path.Join("./config", cf))
		var config map[string]interface{}
		json.Unmarshal(out, &config)
		Enq.ident = append(Enq.ident, identify{ID: config["id"].(uuid.UUID), Name: config["name"].(string)})
	}

}

func (e *Enqueted) Save(interface{}) {
	//verify if alredy exists
	// add in e
	// Save file and return uuid or error
}

func (e *Enqueted) Reload() {
	// zero e
	// init again
}

func (e *Enqueted) Enquete(msg interface{}, header interface{}, id uuid.UUID) {

	for _, i := range e.ident {
		if i.ID == id {

			i.message = append(i.message, message{Msg: msg, Header: header, ID: i.nextID})
			i.nextID++
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
