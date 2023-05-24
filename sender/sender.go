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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func Sender() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer close(interrupt)

	for {
		go func() {
			msg := randString(source.Intn(7950) + 50)
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
			log.Println("stop sending..")
			return
		default:
			time.Sleep(time.Millisecond / 15)
			//time.Sleep(time.Second)
		}
	}
}

func randString(n int) []byte {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}
