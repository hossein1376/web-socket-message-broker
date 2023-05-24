package broker

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"socket/pkg"
)

var (
	receiver    *websocket.Conn
	destination *websocket.Conn
	ch          = make(chan []byte)
)

func Broker() {
	defer receiver.Close()
	defer destination.Close()

	http.HandleFunc("/receiver", receiverHandler)
	http.HandleFunc("/destination", destinationHandler)

	http.ListenAndServe(":3001", nil)
}

func receiverHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	receiver, err = pkg.Upgrader.Upgrade(w, r, nil)
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
	destination, err = pkg.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade error:", err)
		return
	}
}
