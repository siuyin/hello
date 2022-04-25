package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	nc *nats.Conn
	js nats.JetStreamContext

	err error
)

func main() {
	streamCreate()
	clockDisp()
	ticker()

	select {} // wait forever
}

func init() {
	nc, err = nats.Connect("localhost")
	if err != nil {
		log.Fatal("connect: ", err)
	}

	js, err = nc.JetStream()
	if err != nil {
		log.Fatal("jetstream: ", err)
	}
}

func streamCreate() {
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "TIME",
		Subjects: []string{"TIME.*"},
	})
	if err != nil {
		log.Fatal("streamcreate: ", err)
	}
}

func clockDisp() {
	go func() {
		_, err := js.Subscribe("TIME.>", func(m *nats.Msg) {
			fmt.Printf("%s\n", m.Data)
		}, nats.Durable("clockdisp"), nats.DeliverAll())
		if err != nil {
			log.Fatal("replaymsgs: ", err)
		}

		select {} // wait forever
	}()
}

func ticker() {
	go func() {
		for {
			js.Publish("TIME.tick", []byte(fmt.Sprintf("%s", time.Now())))
			time.Sleep(time.Second)
		}
	}()
}
