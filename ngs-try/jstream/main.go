package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

func main() {
	nc := connNGS()
	defer nc.Close()

	js := createStream(nc, "foo")

	pub(js, "foo.a")
	sub(js, "foo.a")

	select {} // wait forever
}

func connNGS() *nats.Conn {
	creds := dflt.EnvString("CREDS", "/h/.nkeys/creds/synadia/SiuYin/SiuYin.creds")
	opt := nats.UserCredentials(creds)
	natsURL := dflt.EnvString("NATS_URL", "nats://connect.ngs.global:4222")
	c, err := nats.Connect(natsURL, opt)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// createStream with retention policy.
func createStream(nc *nats.Conn, name string) nats.JetStreamContext {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := js.AddStream(&nats.StreamConfig{Name: "foo",
		Subjects: []string{"foo.>"}, MaxBytes: 10000000,
		Retention: nats.WorkQueuePolicy}); err != nil {
		log.Fatal(err)
	}

	return js
}

func pub(js nats.JetStreamContext, subj string) {
	go func() {
		for {
			js.Publish(subj, []byte(time.Now().Format("15:04:05.000000")))
			time.Sleep(time.Second)
		}
	}()
}

func sub(js nats.JetStreamContext, subj string) {
	go func() {
		js.Subscribe(subj, func(msg *nats.Msg) {
			fmt.Printf("%s\n", msg.Data)
		}, nats.Durable("fooAConsumer"))
		select {}
	}()
}
