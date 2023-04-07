// Package strm provide high level nats jetstream functions.
package strm

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/siuyin/dflt"
)

// PriceRec is a pricing key-value database record.
type PriceRec struct {
	SKU         string
	Description string
	Price       float64
}

// Server encapsulates an embedded nats jetstream server.
type Server struct {
	svr *server.Server
	nc  *nats.Conn
	js  nats.JetStreamContext
}

// DB is a pricing key-value database.
type DB struct {
	kv nats.KeyValue
}

var s *Server

func svrInit() *Server {
	host := dflt.EnvString("NATS_HOST", "localhost")
	s = &Server{}
	s.svr = newEmbeddedNATSServer(host)
	s.nc = newNATSConn(host)
	s.js = newJetStream(s.nc)
	return s
}

// DBInit sets up a pricing database.
func DBInit(name string) *DB {
	if s == nil {
		s = svrInit()
	}

	db := &DB{}
	db.kv = newKeyValueStore(s.js, name)
	return db
}

// Close closes the pricing database.
func (db *DB) Close() {
	s.nc.Close()
}

// Load loads a csv data stream into the pricing database.
func (db *DB) Load(r io.Reader) error {
	cr := csv.NewReader(r)
	firstRec := true
	for rec, err := cr.Read(); rec != nil && err != io.EOF; rec, err = cr.Read() {
		if firstRec {
			firstRec = false
			continue // skip over one-line header
		}

		price, err := strconv.ParseFloat(rec[2], 64)
		if err != nil {
			return fmt.Errorf("%v: price field does not contain a number", rec)
		}
		val := PriceRec{SKU: rec[0], Description: rec[1], Price: price}
		b, err := json.Marshal(val)
		if err != nil {
			return fmt.Errorf("error marshaling %v: %v", val, err)
		}
		if _, err := db.kv.Put(rec[0], b); err != nil {
			return fmt.Errorf("error writing record: %v: %v", rec[0], err)
		}
	}
	return nil
}

// Dump exports the pricing database records into a csv stream.
func (db *DB) Dump(w io.Writer) error {
	keys, err := db.kv.Keys()
	if err != nil {
		return fmt.Errorf("Unable to list keys: %v", err)
	}
	cw := csv.NewWriter(w)
	if err := cw.Write([]string{"SKU", "Description", "Price"}); err != nil {
		return fmt.Errorf("unable to write csv header: %v", err)
	}
	for i := 0; i < len(keys); i++ {
		v, err := db.kv.Get(keys[i])
		if err != nil {
			return fmt.Errorf("Unable to get key: %v: %v", keys[i], err)
		}
		rec := PriceRec{}
		if err := json.Unmarshal(v.Value(), &rec); err != nil {
			return fmt.Errorf("Unmarshal of %q failed: %v", string(v.Value()), err)
		}
		if err := cw.Write([]string{rec.SKU, rec.Description, fmt.Sprintf("%.2f", rec.Price)}); err != nil {
			return fmt.Errorf("unable to write csv record: %#v: %v", rec, err)
		}
	}
	cw.Flush()

	return nil
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
