package main

import (
	log "github.com/sirupsen/logrus"

	comms "nats-test/example3/good/comms"
)

func main() {

	var c comms.SubscriberComms
	err := c.Init()
	if err != nil {
		panic(err)
	}
	defer c.Ec.Close()

	defer c.Ec.Close()

	log.Info("Connected to NATS and ready to receive messages")

	for {
		msg := <-c.RequestChanRecv
		log.Infof("Received request ID: %d", msg.ID)
	}
}
