package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/ocagent"
	"github.com/siuyin/dflt"
	"github.com/siuyin/hello/world"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"
)

func main() {
	// zPages for debugging
	zPagesMux := http.NewServeMux()
	zpages.Handle(zPagesMux, "/debug")
	go func() {
		if err := http.ListenAndServe(":9999", zPagesMux); err != nil {
			log.Fatalf("Failed to serve zPages")
		}
	}()

	// oce: open census exporter agent
	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithReconnectionPeriod(5*time.Second),
		ocagent.WithAddress("192.168.1.68:55678"), // Only included here for demo purposes.
		ocagent.WithServiceName("helloweb"))
	if err != nil {
		log.Fatalf("Failed to create ocagent-exporter: %v", err)
	}
	trace.RegisterExporter(oce)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// main handler
	subject := dflt.EnvString("SUBJECT", "Siu Yin")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(context.Background(), fmt.Sprintf("%s", r.URL.Path))
		log.Println(r.URL.Path)
		defer span.End()
		fmt.Fprintf(w, "%s %s.\n", world.Greet(), subject)
		goodbye(ctx, w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func goodbye(ctx context.Context, w http.ResponseWriter) {
	_, span := trace.StartSpan(ctx, "goodbye")
	defer span.End()
	fmt.Fprintf(w, "%s.\n", world.GoodBye())
}
