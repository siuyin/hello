package db

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

type DB struct {
	svr *server.Server
	nc  *nats.Conn
	js  nats.JetStreamContext
	kv  nats.KeyValue
}

// Init sets up a pricing database.
func Init(name string) *DB {
	host := dflt.EnvString("NATS_HOST", "localhost")
	db := &DB{}
	db.svr = newEmbeddedNATSServer(host)
	db.nc = newNATSConn(host)
	db.js = newJetStream(db.nc)
	db.kv = newKeyValueStore(db.js, name)
	return db
}

func (db *DB) Close() {
	db.nc.Close()
}

func newEmbeddedNATSServer(host string) *server.Server {
	svr, err := server.NewServer(&server.Options{
		ServerName: "Pricing",
		Host:       host,
		JetStream:  true,
		StoreDir:   "/tmp/pricing",
	})
	if err != nil {
		log.Fatal(err)
	}

	svr.Start()
	for {
		if svr.ReadyForConnections(100 * time.Millisecond) {
			break
		}
	}
	return svr
}

func newNATSConn(host string) *nats.Conn {
	nc, err := nats.Connect(fmt.Sprintf("nats://%s:4222", host))
	if err != nil {
		log.Fatal(err)
	}
	return nc
}

func newJetStream(nc *nats.Conn) nats.JetStreamContext {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	return js
}

func newKeyValueStore(js nats.JetStreamContext, name string) nats.KeyValue {
	kv, err := js.CreateKeyValue(&nats.KeyValueConfig{Bucket: name})
	if err != nil {
		log.Fatal(err)
	}
	return kv
}
