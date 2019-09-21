// example1/natstest-api.go

package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
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

	log.Info("Connected to NATS and ready to send messages")

	personChanSend := make(chan string)
	ec.BindSendChan("hello_subject", personChanSend)

	i := 0
	for {
		time.Sleep(time.Second * 1)
		msg := fmt.Sprintf("Hello World! This is message %d", i)
		log.Infof("Sending %s", msg)
		personChanSend <- msg
		i = i + 1
	}
}
