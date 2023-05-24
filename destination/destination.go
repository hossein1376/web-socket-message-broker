package destination

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func Destination() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3001",
		Path:   "/socket",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("broker failed: %s", err)
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading from socket:", err)
		}
		log.Printf("destination: %s", msg)
	}
}
