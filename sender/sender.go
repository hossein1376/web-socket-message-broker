package sender

import (
	"log"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var source = rand.New(rand.NewSource(time.Now().UnixNano()))

type Socket struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func Sender() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3000",
		Path:   "/",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("sender failed: %s", err)
	}
	defer c.Close()

	socket := &Socket{
		conn: c,
	}

	for {
		go worker(socket)
		time.Sleep(time.Millisecond / 100)
	}
}

func worker(s *Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := randString(source.Int31n(8193))

	err := s.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func randString(n int32) []byte {
	b := make([]byte, n+2)
	source.Read(b)
	return b
}
