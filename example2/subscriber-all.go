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

	// Make sure this type and its properties are exported
	// so the serializer doesn't bork
	type Request struct {
		ID int
	}
	requestChanRecv := make(chan *Request)
	ec.BindRecvChan("request_subject", requestChanRecv)

	for {
		// Wait for incoming messages
		req := <-requestChanRecv

		log.Infof("Received request: %d", req.ID)
	}
}
