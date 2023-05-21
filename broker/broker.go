package broker

import (
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Socket struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func Broker() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3002",
		Path:   "/",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("broker failed: %s", err)
	}
	defer conn.Close()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("broker upgrade failed: %s", err)
		}
		defer c.Close()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}

			socket := &Socket{
				conn: c,
			}

			go worker(socket, msg)
		}
	})

	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func worker(s *Socket, msg []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Fatal(err)
	}
}
