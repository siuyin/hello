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
	receiver()
}

func init() {
	nc, err = nats.Connect("nats://localhost:4222") // note connection to port 4222
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

func receiver() {
	sub, err := js.PullSubscribe("posts.*", "clockdisp", nats.Bind("posts", "clockdisp"))
	if err != nil {
		log.Fatal("pullsub: ", err)
	}

	for {
		fmt.Println("Running receiver")
		//for msgs, err := sub.Fetch(10, nats.MaxWait(2*time.Second)); len(msgs) > 0 && err == nil; msgs, err = sub.Fetch(10, nats.MaxWait(2*time.Second)) {
		msgs, err := sub.Fetch(10, nats.MaxWait(20*time.Second))
		if err != nil {
			log.Fatal("fetch: ", err)
		}
		fmt.Println("num msgs: ", len(msgs))
		for _, msg := range msgs {
			fmt.Printf("%s\n", msg.Data)
			msg.Ack()
		}
	}
}
