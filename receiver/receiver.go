package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
