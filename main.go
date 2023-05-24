package main

import (
	"time"

	"socket/broker"
	"socket/destination"
	"socket/receiver"
	"socket/sender"
)

func main() {
	go broker.Broker()
	go receiver.Receiver()
	go destination.Destination()
	go sender.Sender()

	time.Sleep(time.Second * 10)
}
