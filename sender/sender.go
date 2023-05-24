package sender

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var source = rand.New(rand.NewSource(time.Now().UnixNano()))

func Sender() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer close(interrupt)

	for {
		go func() {
			msg := randString(5)
			body := bytes.NewReader(msg)

			req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/", body)
			if err != nil {
				log.Fatalf("make http request: %v", err)
				return
			}

			_, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Fatalf("send http request: %v", err)
				return
			}
		}()

		select {
		case <-interrupt:
			log.Println("interrupt")
			return
		default:
			time.Sleep(time.Millisecond / 20)
			//time.Sleep(time.Second)
		}
	}
}

func randString(n int32) []byte {
	b := make([]byte, n+2)
	source.Read(b)
	return b
}
