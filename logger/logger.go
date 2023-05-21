package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3000",
		Path:   "/",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("logger dialer:", err)
		return
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("logger:", err)
				return
			}
			log.Printf("logger recvied: %s", message)
		}
	}()

	for {

	}
}
