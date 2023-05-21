package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"

	"socket/pkg"
)

func main() {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:3000",
		Path:   "/",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("sender dial:", err)
		return
	}
	defer c.Close()

	client := &pkg.Client{
		Conn: c,
	}

	msg := pkg.Message{
		Target: pkg.Receiver,
		Body:   nil,
	}

	for {
		go worker(client, msg)
		//time.Sleep(time.Millisecond / 100)
		time.Sleep(time.Second)
	}
}

func worker(c *pkg.Client, msg pkg.Message) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	msg.Body = randString(pkg.Source.Int31n(8193))

	// TODO: parse it to json and then send it

	err := c.Conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func randString(n int32) []byte {
	b := make([]byte, n+2)
	pkg.Source.Read(b)
	return b
}
