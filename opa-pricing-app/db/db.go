// Package strm provide high level nats jetstream functions.
package db

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/nats-io/nats.go"
	"github.com/siuyin/strm"
)

// PriceRec is a pricing key-value database record.
type PriceRec struct {
	SKU         string
	Description string
	Price       float64
}
type DB struct {
	s *strm.DB
}

// Init sets up a pricing database.
func Init(name string) *DB {
	db := DB{}
	db.s = strm.DBInit(name)
	return &db
}

// Put stores a value with the specified key.
func (db *DB) Put(key string, value []byte) (uint64, error) {
	return db.s.KV.Put(key, value)
}

// Get retrieves the value stored at the specified key.
func (db *DB) Get(key string) (nats.KeyValueEntry, error) {
	return db.s.KV.Get(key)
}

// Delete delete the entry at key.
func (db *DB) Delete(key string) error {
	return db.s.KV.Delete(key)
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
		if _, err := db.s.KV.Put(rec[0], b); err != nil {
			return fmt.Errorf("error writing record: %v: %v", rec[0], err)
		}
	}
	return nil
}

// Dump exports the pricing database records into a csv stream.
func (db *DB) Dump(w io.Writer) error {
	keys, err := db.s.KV.Keys()
	if err != nil {
		return fmt.Errorf("Unable to list keys: %v", err)
	}
	cw := csv.NewWriter(w)
	if err := cw.Write([]string{"SKU", "Description", "Price"}); err != nil {
		return fmt.Errorf("unable to write csv header: %v", err)
	}
	for i := 0; i < len(keys); i++ {
		v, err := db.s.KV.Get(keys[i])
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

func (db *DB) Close() {
	db.s.Close()
}
