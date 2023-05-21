package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Content string
}

type Stats struct {
	Count int
	Size  int
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	go sendMessages()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Upgrade error:", err)
		return
	}
	defer conn.Close()

	receiveMessages(conn)
}

func sendMessages() {
	for {
		time.Sleep(1 * time.Second)

		message := Message{
			Content: fmt.Sprintf("Timestamp: %s", time.Now().Format(time.RFC3339)),
		}

		sendMessageToClients(message)
	}
}

var connections []*websocket.Conn
var connectionsMutex sync.Mutex

func sendMessageToClients(message Message) {
	connectionsMutex.Lock()
	defer connectionsMutex.Unlock()

	for _, conn := range connections {
		err := conn.WriteJSON(message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			continue
		}
	}
}

func receiveMessages(conn *websocket.Conn) {
	connectionsMutex.Lock()
	connections = append(connections, conn)
	connectionsMutex.Unlock()

	defer func() {
		connectionsMutex.Lock()
		connections = removeConnection(connections, conn)
		connectionsMutex.Unlock()
	}()

	stats := Stats{
		Count: 0,
		Size:  0,
	}

	for {
		message := Message{}

		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		stats = logMessage(message, stats)
		displayStats(stats)
	}
}

func logMessage(message Message, stats Stats) Stats {
	log.Printf("Received message: %s", message.Content)
	stats.Count++
	stats.Size += len(message.Content)
	return stats
}

func displayStats(stats Stats) {
	log.Printf("Total messages: %d, total size: %d", stats.Count, stats.Size)
}

func removeConnection(slice []*websocket.Conn, item *websocket.Conn) []*websocket.Conn {
	for i, conn := range slice {
		if conn == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
