package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/siuyin/dflt"
)

type hub struct {
	subj string
	addr string
	conn stan.Conn
}

func (h *hub) connect(clusterID, clientID string) error {
	var err error
	h.conn, err = stan.Connect(clusterID, clientID, stan.NatsURL(h.addr))
	if err != nil {
		return err
	}
	return nil
}
func (h *hub) publish(msg string) error {
	h.conn.Publish(h.subj, []byte(msg))
	return nil
}

func main() {
	fmt.Println("nats streaming example")
	h := hub{subj: "junk",
		addr: dflt.EnvString("NATS_URL", "nats://192.168.1.68:4222")}
	if err := h.connect(dflt.EnvString("CLUSTER_ID", "test-cluster"), "junk-client"); err != nil {
		log.Fatal(err)
	}
	h.conn.Subscribe(h.subj, func(m *stan.Msg) {
		fmt.Printf("seq: %v, data: %s\n", m.Sequence, m.Data)
	}, stan.DurableName("junk"))
	for {
		h.publish(time.Now().Format("15:04:05.000"))
		time.Sleep(time.Second)
	}
}
