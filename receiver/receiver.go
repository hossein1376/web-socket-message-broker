package receiver

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	ch = make(chan []byte)
)

func Receiver() {
	http.HandleFunc("/", apiHandler)

	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal("failed to start receiver server: ", err)
		}
	}()

	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3001",
		Path:   "/receiver",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("broker failed: %s", err)
	}
	defer conn.Close()

	for {
		// receiving messages from the apiHandler using a channel; and sending them to the broker
		err = conn.WriteMessage(websocket.TextMessage, <-ch)
		if err != nil {
			log.Println("error writing to broker socket:", err)
		}
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// sending the message on a channel, so it can be received on the other handler
	ch <- msg
}
