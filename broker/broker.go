package broker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	receiver    *websocket.Conn
	destination *websocket.Conn
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func Broker() {
	http.HandleFunc("/receiver", receiverHandler)
	http.HandleFunc("/destination", destinationHandler)

	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("failed to start broker server:", err)
		return
	}
}

func receiverHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	receiver, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("receiver upgrade:", err)
		return
	}

	for {
		_, msg, err := receiver.ReadMessage()
		if err != nil {
			log.Println("error reading:", err)
			return
		}

		//log.Print("logger: ", string(msg))

		err = destination.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("error writing to destination:", err)
		}
	}
}

func destinationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	destination, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
}
