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
	replayMsgs("LOGGER")
	pubAsyncMsgs(54)
	syncPubMsgs(33)
	//consume("MONITOR")
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
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	if err != nil {
		log.Fatal("streamcreate: ", err)
	}
}

func pubAsyncMsgs(n int) {
	go func() {
		for i := 0; i < n; i++ {
			js.PublishAsync("ORDERS.scratch", []byte(fmt.Sprintf("hello: %v", i)))
		}
		//select {
		//case <-js.PublishAsyncComplete():
		//	fmt.Println("Done")
		//case <-time.After(5 * time.Second):
		//	fmt.Println("Did not resolve in time")
		//}
		<-js.PublishAsyncComplete()
	}()
}

func syncPubMsgs(n int) {
	for i := 0; i < n; i++ {
		js.Publish("ORDERS.sync", []byte(fmt.Sprintf("sync: %v", i)))
	}
}

func consume(consumer string) {
	sub, err := js.PullSubscribe("ORDERS.*", consumer)
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

func replayMsgs(consumer string) {
	go func() {
		_, err := js.Subscribe("ORDERS.>", func(m *nats.Msg) {
			fmt.Printf("%s\n", m.Data)
		}, nats.Durable(consumer), nats.DeliverAll())
		if err != nil {
			log.Fatal("replaymsgs: ", err)
		}

		select {} // wait forever
	}()
}
