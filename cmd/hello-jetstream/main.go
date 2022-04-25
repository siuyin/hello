package main

import (
	"context"
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
	consume("MONITOR")
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("all done, please ignore the context deadline exceeded message")
			return
		default:
		}

		msgs, err := sub.Fetch(100, nats.Context(ctx))
		if err != nil {
			log.Fatal("fetch: ", err)
		}
<<<<<<< HEAD
=======
		fmt.Println("fetching ", len(msgs), " messages")
>>>>>>> 75341578c6b0ee4bd058367e2697ae9c17772479
		for _, msg := range msgs {
			fmt.Printf("%s\n", msg.Data)
			msg.Ack()
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
