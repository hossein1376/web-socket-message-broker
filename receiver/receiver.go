package receiver

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Socket struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

//func main() {
//	u := url.URL{
//		Scheme: "ws",
//		Host:   "localhost:3000",
//		Path:   "/",
//	}
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer c.Close()
//	var counter int32
//	var t = time.Now()
//
//	for {
//		_, msg, err := c.ReadMessage()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		counter++
//		if counter%10000 == 0 {
//			log.Println(counter, time.Since(t))
//			fmt.Println(len(msg))
//			t = time.Now()
//		}
//
//
//	}
//}

func Receiver() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3001",
		Path:   "/",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("receiver failed: %s", err)
	}
	defer conn.Close()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("Failed to upgrade WebSocket connection: %s", err)
		}
		defer c.Close()

		var counter int32
		var t = time.Now()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			counter++
			if counter%10000 == 0 {
				log.Println(counter, time.Since(t))
				fmt.Println(len(msg))
				t = time.Now()
			}

			socket := &Socket{
				conn: c,
			}

			go worker(socket, msg)
		}
	})

	err = http.ListenAndServe(":3000", nil)
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
