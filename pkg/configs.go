package pkg

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Target int8

const (
	Sender Target = iota
	Receiver
	Logger
	Destination
)

var (
	Upgrader = websocket.Upgrader{}
	Source   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Message struct {
	Target Target
	Body   []byte
}

type Client struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}
