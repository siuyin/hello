package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

func main() {
	fmt.Println("connecting to NGS")

	nc := connNGS()
	defer nc.Close()

	pub(nc)
	sub(nc)

	select {}
}

func connNGS() *nats.Conn {
	creds := dflt.EnvString("CREDS", "/h/.nkeys/creds/synadia/SiuYin/SiuYin.creds")
	opt := nats.UserCredentials(creds)
	//natsURL := dflt.EnvString("NATS_URL", "nats://connect.ngs.global:4222")
	natsURL := dflt.EnvString("NATS_URL", "connect.ngs.global")
	c, err := nats.Connect(natsURL, opt)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

const subj = "test"

func pub(nc *nats.Conn) {
	go func() {
		for {
			burst(nc, 3)
			time.Sleep(3 * time.Second)
		}
	}()
}
func burst(nc *nats.Conn, n int) {
	for i := 0; i < n; i++ {
		tm := time.Now().Format("15:04:05.000000")
		nc.Publish("test.a.b", []byte(tm))
	}
}

func sub(nc *nats.Conn) {
	nc.Subscribe("test.>", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s: len: %d\n", string(m.Data), len(m.Data))
		time.Sleep(900 * time.Millisecond)
	})
}
