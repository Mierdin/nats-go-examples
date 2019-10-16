package comms

import (
	nats "github.com/nats-io/nats.go"
)

const (
	SUBJECT_HELLO = "hello_subject"
)

func NewPublisherComms() (*PublisherComms, error) {

	pc := PublisherComms{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	pc.Nc = nc

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	pc.Ec = ec

	// Bind send channel
	pc.RequestChanSend = make(chan *Request)
	pc.Ec.BindSendChan(SUBJECT_HELLO, pc.RequestChanSend)

	return &pc, nil
}

func NewSubscriberComms() (*SubscriberComms, error) {

	sc := SubscriberComms{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	sc.Nc = nc

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		return nil, err
	}
	sc.Ec = ec

	// Bind recieve channel using Queues
	sc.RequestChanRecv = make(chan *Request)
	sc.Ec.BindRecvQueueChan(SUBJECT_HELLO, "hello_queue", sc.RequestChanRecv)

	return &sc, nil
}
