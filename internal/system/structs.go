package system

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

/*System - Main struct for system*/
type System struct {
	config     Conf
	Queue      []*queueConf
	Dispatcher []interface{}
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

/*queue internal configuration*/
type queueConf struct {
	Name       string
	ID         uuid.UUID
	persistent bool
	TTL        time.Duration
	maxSize    int64
	createDate time.Time
	NextID     int64
	Variable   []Variable
	Messages   []*message
	mutex      sync.Mutex
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
