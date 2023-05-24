package destination

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Stat struct {
	count int32
	size  int
}

func Destination() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3001",
		Path:   "/destination",
	}

	stat := Stat{}

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

		stat.count++
		stat.size += len(msg)

		//log.Printf("destination: %s", msg)
		if stat.count%10000 == 0 {
			log.Printf("destination: %d messages, %d bytes", stat.count, stat.size)
		}
	}
}
