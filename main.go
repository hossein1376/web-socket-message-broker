package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"socket/broker"
	"socket/destination"
	"socket/receiver"
	"socket/sender"
)

func main() {
	// make interrupt channel and listen for the os signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer close(interrupt)

	start := time.Now()
	stat := destination.Stat{}

	// start all the goroutines
	go sender.Sender()
	go receiver.Receiver()
	go broker.Broker()
	go destination.Destination(&stat)

	select {
	// wait for the interrupt signal, then print the stats and exit
	case <-interrupt:
		fmt.Printf(
			"\nrecived %d messages, %d kilo-bytes, %d mega-bytes in %f seconds\n",
			stat.Count,
			stat.Size/1024,
			stat.Size/(1024*1024),
			time.Since(start).Seconds())
		return
	}
}
