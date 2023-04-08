package db

import (
	"bytes"
	"encoding/csv"
	"os"
	"testing"

	"github.com/nats-io/nuid"
)

func TestDBFunctions(t *testing.T) {
	bucketName := nuid.Next()
	db := Init(bucketName)

	t.Run("put and get", func(t *testing.T) {
		if _, err := db.Put("a", []byte("apple")); err != nil {
			t.Error(err)
		}
		if v, err := db.Get("a"); string(v.Value()) != "apple" || err != nil {
			t.Errorf("unable to get key 'a':  %v", err)
		}
		if err := db.Delete("a"); err != nil {
			t.Error(err)
		}
	})

	t.Run("load and dump", func(t *testing.T) {
		f, err := os.Open("testdata/sample.csv")
		if err != nil {
			t.Error(err)
		}
		if err := db.Load(f); err != nil {
			t.Error(err)
		}

		var buf bytes.Buffer
		if err := db.Dump(&buf); err != nil {
			t.Error(err)
		}
		cr := csv.NewReader(&buf)
		recs, err := cr.ReadAll()
		if err != nil {
			t.Error(err)
		}
		if len(recs) != 3 {
			t.Error("Unexpected number of records")
		}

	})
	if err := db.s.Delete(bucketName); err != nil {
		t.Error(err)
	}
	db.Close()
}
