package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("localhost")
	if err != nil {
		log.Fatal("connect: ", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("jetstream: ", err)
	}

	// Create a Stream
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	// Simple Async Stream Publisher
	for i := 0; i < 504; i++ {
		js.PublishAsync("ORDERS.scratch", []byte(fmt.Sprintf("hello: %v", i)))
	}
	//select {
	//case <-js.PublishAsyncComplete():
	//	fmt.Println("Done")
	//case <-time.After(5 * time.Second):
	//	fmt.Println("Did not resolve in time")
	//}
	<-js.PublishAsyncComplete()

	for i := 0; i < 303; i++ {
		js.Publish("ORDERS.sync", []byte(fmt.Sprintf("sync: %v", i)))
	}

	// Simple Pull Consumer
	sub, err := js.PullSubscribe("ORDERS.*", "MONITOR")
	if err != nil {
		log.Fatal("pullsub: ", err)
	}

	for msgs, err := sub.Fetch(100, nats.MaxWait(1*time.Second)); len(msgs) > 0 && err == nil; msgs, err = sub.Fetch(100, nats.MaxWait(1*time.Second)) {
		fmt.Println("fetching ", len(msgs), " messages")
		if err != nil {
			log.Fatal("fetch: ", err)
		}
		for i := 0; i < len(msgs); i++ {
			fmt.Printf("%s\n", msgs[i].Data)
		}
	}
}
