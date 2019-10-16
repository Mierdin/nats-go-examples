package comms

import (
	nats "github.com/nats-io/nats.go"
)

const (
	SUBJECT_HELLO = "hello_subject"
)

// Request is the Go type that we will be using to pass messages between
// our services.
type Request struct {
	ID int
}

// Comms centralizes connection objects and channels so we have one place to
// go in order to send or receive messages with NATS.
type Comms struct {

	// NATS connection types
	Nc *nats.Conn
	Ec *nats.EncodedConn

	// Send/Receive Channels
	RequestChanSend chan *Request
	RequestChanRecv chan *Request
}

// NewComms returns an instance of Comms with a running connection to the NATS server
// and channels pre-bound to NATS subjects, ready to send/receive messages
func NewComms() (*Comms, error) {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}

	cc := Comms{}

	// Bind recieve channel using Queues
	cc.RequestChanRecv = make(chan *Request)
	ec.BindRecvQueueChan(SUBJECT_HELLO, "hello_queue", cc.RequestChanRecv)

	// Bind send channel
	cc.RequestChanSend = make(chan *Request)
	ec.BindSendChan(SUBJECT_HELLO, cc.RequestChanSend)

	return &cc, nil
}
