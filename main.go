package main

import (
	"time"

	"graph/broker"
	"graph/destination"
	"graph/receiver"
	"graph/sender"
)

func main() {
	go destination.Destination()
	time.Sleep(time.Second)
	go broker.Broker()
	time.Sleep(time.Second)
	go receiver.Receiver()
	time.Sleep(time.Second)
	go sender.Sender()

	time.Sleep(10 * time.Second)
}
