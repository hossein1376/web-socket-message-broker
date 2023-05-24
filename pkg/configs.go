package pkg

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Target int8

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	Source = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Message struct {
	Target Target
	Body   []byte
}

type Client struct {
	Conn  *websocket.Conn
	Mutex sync.Mutex
}
