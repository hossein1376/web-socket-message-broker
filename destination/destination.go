package destination

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type Stat struct {
	Count int
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

	// timer every second
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		// read messages from broker
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("error reading from broker socket:", err)
		}

		// update the stats
		stat.Count++
		stat.Size += len(msg)

		// print stats every one second
		select {
		case <-ticker.C:
			log.Printf(
				"destination: %d messages, %d kilo-bytes, %d mega-bytes",
				stat.Count,
				stat.Size/1024,
				stat.Size/(1024*1024))
		default:
			// do nothing
		}
	}
}
