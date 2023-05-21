package destination

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Destination() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatalf("destination failed: %s", err)
		}
		defer c.Close()

		var counter int32

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}

			counter++
			if counter%30000 == 0 {
				fmt.Printf("%s\n", msg)
			}

		}
	})

	err := http.ListenAndServe(":3002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
