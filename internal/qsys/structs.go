package qsys

import (
	"go-queue/internal/qsys/queue"
	"sync"
)

/*System - Main struct for system*/
type QSys struct {
	config     Conf
	queue      *queue.Queue
	dispatcher []interface{}
	mutex      sync.Mutex
}

/*ConfigQueue - receive configuration from customer*/
type ConfigQueue struct {
	Name       string
	Persistent bool
	TTL        int
	MaxSize    int64
	Variable   []Variable
}

// Variable component config
type Variable struct {
	Key   string
	Value string
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
