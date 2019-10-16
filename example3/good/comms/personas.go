package comms

import (
	nats "github.com/nats-io/nats.go"
)

type Comms struct {
	// NATS connection types
	Nc *nats.Conn
	Ec *nats.EncodedConn
}

type PublisherComms struct {
	Comms

	// Send Channel
	RequestChanSend chan *Request
}

type SubscriberComms struct {
	Comms

	// Receive Channel
	RequestChanRecv chan *Request
}
