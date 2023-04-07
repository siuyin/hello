package strm

import (
	"bytes"
	"encoding/csv"
	"os"
	"testing"

	"github.com/nats-io/nuid"
)

func TestDBFunctions(t *testing.T) {
	bucketName := nuid.Next()
	db := DBInit(bucketName)

	t.Run("put and get", func(t *testing.T) {
		if _, err := db.kv.PutString("a", "apple"); err != nil {
			t.Error(err)
		}
		if v, err := db.kv.Get("a"); string(v.Value()) != "apple" || err != nil {
			t.Errorf("unable to get key 'a':  %v", err)
		}
		if err := db.kv.Delete("a"); err != nil {
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
	if err := s.js.DeleteKeyValue(bucketName); err != nil {
		t.Error(err)
	}
	db.Close()
}
