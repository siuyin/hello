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
	tickAudit()

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

func tickAudit() {
	go func() {
		sub, err := js.PullSubscribe("TIME.tick", "tickauditor")
		if err != nil {
			log.Fatal("pullsub: ", err)
		}

		for {
			for msgs, err := sub.Fetch(100, nats.MaxWait(1*time.Second)); len(msgs) > 0 && err == nil; msgs, err = sub.Fetch(100, nats.MaxWait(1*time.Second)) {
				//fmt.Println("fetching ", len(msgs), " messages")
				if err != nil {
					log.Fatal("fetch: ", err)
				}
				for i := 0; i < len(msgs); i++ {
					fmt.Printf("audit: %s\n", msgs[i].Data)
					msgs[i].Ack()
				}
			}
		}
	}()
}
