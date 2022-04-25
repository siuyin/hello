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
	poster()

	select {} // wait forever
}

func init() {
	nc, err = nats.Connect("nats://localhost:4223") // custom port 4223
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
		Name:     "posts",
		Subjects: []string{"posts.*"},
	})
	if err != nil {
		log.Fatal("streamcreate: ", err)
	}
}

func poster() {
	go func() {
		tkr := time.NewTicker(time.Second)
		defer tkr.Stop()

		for {
			js.Publish("posts.tick", []byte(fmt.Sprintf("%s", time.Now().Format("15:04:05"))))
			<-tkr.C
		}
	}()
}
