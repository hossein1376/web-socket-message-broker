package sender

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var source = rand.New(rand.NewSource(time.Now().UnixNano()))

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	bits = 6
	mask = 1<<bits - 1
	max  = 63 / bits
)

func Sender() {
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

		// for 10000 messages per second, we need to send 1 each 1/10 ms
		// here, we actually sleep 1/15 ms to compensate for the overhead of the sending
		time.Sleep(time.Millisecond / 15)
	}
}

// generate random string using bitmask for high performance
func randString(n int) []byte {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), max; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), max
		}
		if idx := int(cache & mask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= bits
		remain--
	}
	return b
}
