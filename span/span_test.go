package span

import (
	"fmt"
	"testing"

	"go.opencensus.io/trace"
)

func TestWrapAndUnwrap(t *testing.T) {
	b := []byte("Hello")
	sc := trace.SpanContext{}
	copy(sc.TraceID[:], "0123456789012345")
	copy(sc.SpanID[:], "abcdefgh")
	sc.TraceOptions = 456
	o, err := Wrap(b, sc)
	if err != nil {
		t.Error(err)
	}

	b1, sc1, err := Unwrap(o)
	if err != nil {
		t.Error(err)
	}
	if string(b1) != string(b) {
		t.Errorf("Unexpected output: %s", b1)
	}

	s := fmt.Sprintf("%v", sc)
	s1 := fmt.Sprintf("%v", sc1)
	if s != s1 {
		t.Errorf("Unexpected value: %s", s1)
	}
}
