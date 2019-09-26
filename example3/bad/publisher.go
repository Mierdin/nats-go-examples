package main

import (
	"time"

	comms "nats-test/example3/bad/comms"

	log "github.com/sirupsen/logrus"
)

func main() {

	// A single type for all comms means all our communication channels are
	// properties of c.
	c, err := comms.NewComms()
	if err != nil {
		panic(err)
	}
	defer c.Ec.Close()

	log.Info("Connected to NATS and ready to send messages")

	i := 0
	for {

		// Create instance of type Request with Id set to
		// the current value of i
		req := comms.Request{ID: i}

		// Just send to the channel! :)
		log.Infof("Sending request %d", req.ID)
		c.PersonChanSend <- &req

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1
	}
}
