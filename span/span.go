// Package span wraps and unwraps an opencensus trace.SpanContext.
package span

import (
	"encoding/json"
	"fmt"

	"go.opencensus.io/trace"
)

// Data embeds a SpanContext
type Data struct {
	trace.SpanContext
	Data []byte
}

// Wraps wraps a SpanContext
func Wrap(b []byte, sc trace.SpanContext) ([]byte, error) {
	d := new(Data)
	d.TraceID = sc.TraceID
	d.SpanID = sc.SpanID
	d.TraceOptions = sc.TraceOptions
	d.Data = make([]byte, len(b))
	copy(d.Data, b)
	bs, err := json.Marshal(d)
	if err != nil {
		return b, fmt.Errorf("Wrap could not marshal json: %v: %s", err, b)
	}
	return bs, nil
}

// Unwrap unwraps a SpanContext
func Unwrap(b []byte) ([]byte, trace.SpanContext, error) {
	d := new(Data)
	if err := json.Unmarshal(b, d); err != nil {
		return b, trace.SpanContext{}, err
	}
	return d.Data, trace.SpanContext{TraceID: d.TraceID, SpanID: d.SpanID, TraceOptions: d.TraceOptions}, nil
}

func ChildLink(sc trace.SpanContext, msg string) trace.Link {
	return trace.Link{
		TraceID:    sc.TraceID,
		SpanID:     sc.SpanID,
		Type:       trace.LinkTypeChild,
		Attributes: map[string]interface{}{"message": msg},
	}
}
