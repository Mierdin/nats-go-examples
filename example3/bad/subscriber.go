package main

import (
	log "github.com/sirupsen/logrus"

	comms "nats-test/example3/bad/comms"
)

func main() {

	// A single type for all comms means all our communication channels are
	// properties of c.
	c, err := comms.NewComms()
	if err != nil {
		panic(err)
	}
	defer c.Ec.Close()

	log.Info("Connected to NATS and ready to receive messages")

	for {
		msg := <-c.PersonChanRecv
		log.Infof("Received request ID: %d", msg.ID)
	}
}
