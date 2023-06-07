# Web Socket Message Broker

## Description

This is a test project that I did a while back for a job interview, but I liked the challenge and decided to share my
code.

The question was to write a program that sends 10,000 messages per second to a broker which in turn it will receive, log and send
them to the destination. In there, the overall message count and size will be displayed.  
Sender module should use standard HTTP protocol and the broker should use the WebSocket protocol.

## Dependencies

This project uses `gorilla/websocket`, install it with:

```go
go get github.com/gorilla/websocket
```

## Running

The project has four modules that are called from the `main.go` file. Run it using

```go
go run ./main.go
```

## Sender

The `sender` module generates random messages between 50 bytes and 8 kilobytes, and send them to `receiver`'s HTTP
endpoint.  
There is a tiny sleep delay to control the program to ~10k messages per seconds. It can be tweaked for higher or lower
counts.

## Receiver

It starts a HTTP server on port 3000 and listen for incoming messages, which in turn will redirect them to a WebSocket
connection.

## Broker

The `broker` module starts a server on port 3001 and upgrades the handlers to WebSocket.  
It will log the messages to the `broker.log` file in the root directory of the project. Then, the messages will be sent
to the destination module.

## Destination

After receiving each message, it will increment the total count and the received size.  
It will print the stats every one second.

## Stopping

To stop the application, send the interrupt signal (`ctrl + c`). It will display the overall statics and then will exit.