// example1/natstest-worker.go

package main

import (
	nats "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func main() {

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	log.Info("Connected to NATS and ready to receive messages")

	type Request struct {
		ID int
	}
	requestChanRecv := make(chan *Request)

	// This allows us to subscribe to a queue within a subject
	// for load balancing messages among subscribers.
	// https://godoc.org/github.com/nats-io/go-nats#EncodedConn.BindRecvQueueChan
	ec.BindRecvQueueChan("request_subject", "hello_queue", requestChanRecv)

	for {
		// Wait for incoming messages
		req := <-requestChanRecv
		log.Infof("Received request: %d", req.ID)
	}
}
