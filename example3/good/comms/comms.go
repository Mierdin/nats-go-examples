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

// Commser gives us a way of describing how all communcations-related "personas"
// should behave. They should all have a Init() function that returns an instance of them
// with the right setup (only the necessary channels activated)
type Commser interface {
	Init()
}

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

func (pc *PublisherComms) Init() error {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	// Bind send channel
	pc.RequestChanSend = make(chan *Request)
	ec.BindSendChan(SUBJECT_HELLO, pc.RequestChanSend)

	return nil
}

type SubscriberComms struct {
	Comms

	// Receive Channel
	RequestChanRecv chan *Request
}

func (sc *SubscriberComms) Init() error {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return err
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return err
	}

	// Bind recieve channel using Queues
	sc.RequestChanRecv = make(chan *Request)
	ec.BindRecvQueueChan(SUBJECT_HELLO, "hello_queue", sc.RequestChanRecv)

	return nil
}
