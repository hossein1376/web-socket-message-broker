package broker

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var (
	logger      *log.Logger
	receiver    *websocket.Conn
	destination *websocket.Conn
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func Broker() {
	// open log file
	output, err := os.OpenFile("./broker.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("failed to open log file:", err)
	}
	defer output.Close()

	// instantiate the logger
	logger = log.New(output, "broker: ", log.Ltime|log.Lmsgprefix)

	http.HandleFunc("/receiver", receiverHandler)
	http.HandleFunc("/destination", destinationHandler)

	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("failed to start broker server:", err)
		return
	}
}

func receiverHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	receiver, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("failed receiver upgrade:", err)
		return
	}

	for {
		// receive the message from the receiver
		_, msg, err := receiver.ReadMessage()
		if err != nil {
			log.Println("error reading from receiver socket:", err)
			return
		}

		// print the message to the logger
		logger.Print(string(msg))

		// send the message to the destination
		err = destination.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("error writing to destination socket:", err)
		}
	}
}

// destinationHandler establishes a WebSocket connection with the destination module
func destinationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	destination, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("failed destination error:", err)
		return
	}
}
