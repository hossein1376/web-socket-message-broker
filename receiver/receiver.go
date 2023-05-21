package main

import (
	"log"
	"net/http"

	"socket/pkg"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := pkg.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("receiver upgrade:", err)
			return
		}

		defer c.Close()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("receiver:", err)
				break
			}
			log.Printf("recvied: %s", msg)
		}
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
