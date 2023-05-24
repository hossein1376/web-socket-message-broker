package destination

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Stat struct {
	Count int32
	Size  int
}

func Destination(stat *Stat) {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3001",
		Path:   "/destination",
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

		stat.Count++
		stat.Size += len(msg)

		//log.Printf("destination: %s", msg)
		//if stat.Count%10000 == 0 {
		//	log.Printf("destination: %d messages, %d kilo-bytes, %d mega-bytes", stat.Count, stat.Size/1024, stat.Size/(1024*1024))
		//}
	}
}
