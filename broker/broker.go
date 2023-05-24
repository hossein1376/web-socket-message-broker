package broker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"socket/pkg"
)

var ch = make(chan []byte)

func Broker() {
	http.HandleFunc("/socket", socketHandler)

	http.ListenAndServe(":3001", nil)

}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := pkg.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("receiver upgrade:", err)
		return
	}
	defer c.Close()

	go func() {
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("error reading:", err)
				return
			}

			log.Print("logger: ", string(msg))

			ch <- msg
		}
	}()

	for {
		err = c.WriteMessage(websocket.TextMessage, <-ch)
		if err != nil {
			log.Println("error writing to socket:", err)
		}
	}

}
