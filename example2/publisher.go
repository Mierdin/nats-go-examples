// example1/natstest-api.go

package main

import (
	"time"

	"github.com/nats-io/nats.go"
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

	log.Info("Connected to NATS and ready to send messages")

	type Request struct {
		ID int
	}
	requestChanSend := make(chan *Request)
	ec.BindSendChan("request_subject", requestChanSend)

	i := 0
	for {

		// Create instance of type Request with ID set to
		// the current value of i
		req := Request{ID: i}

		// Just send to the channel! :)
		log.Infof("Sending request %d", req.ID)
		requestChanSend <- &req

		// Pause and increment counter
		time.Sleep(time.Second * 1)
		i = i + 1
	}
}
