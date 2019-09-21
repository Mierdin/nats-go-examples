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

	ec, err := nats.NewEncodedConn(nc, nats.DEFAULT_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	log.Info("Connected to NATS and ready to receive messages")

	personChanRecv := make(chan string)
	ec.BindRecvChan("hello_subject", personChanRecv)

	for {
		msg := <-personChanRecv

		log.Infof("Received: %s", msg)
	}
}
