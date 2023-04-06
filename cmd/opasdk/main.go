package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/open-policy-agent/opa/sdk"
	svr "github.com/open-policy-agent/opa/sdk"
	sdktest "github.com/open-policy-agent/opa/sdk/test"
)

func main() {
	ctx := context.Background()

	// create a mock HTTP bundle server
	server, err := sdktest.NewServer(sdktest.MockBundle("/bundles/bundle.tar.gz", map[string]string{
		"example.rego": `
				package authz
				import future.keywords

				default allow := false
				allow if input.open == data.smallSesame

				default value := 0.0
				value := 10.0 if { 
				input.open == data.bigSesame
				}
				value := -5.0 if { 
				input.open == data.smallSesame
				}
			`,
		"data.json": `{
"smallSesame": "sesame",
"bigSesame": "Sesame"
}`,
	}))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bundle server at:", server.URL())

	defer server.Stop()

	config := []byte(fmt.Sprintf(`{
"services": {
  "test": { "url": %q }
},
"bundles": {
  "test": {"resource": "/bundles/bundle.tar.gz"}
},
"decision_logs": {"console":false}
}
`, server.URL()))

	opa, err := svr.New(ctx, sdk.Options{
		Config: bytes.NewReader(config),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer opa.Stop(ctx)

	start := time.Now()
	n := 5
	for i := 0; i < n; i++ {
		q := "/authz/value"
		inp := map[string]interface{}{"open": "Sesame"}
		result := decide(opa, q, inp)

		fmt.Println(result.Result)
	}
	for i := 0; i < n; i++ {
		q := "/authz/allow"
		inp := map[string]interface{}{"open": "sesame"}
		decide(opa, q, inp)
		result := decide(opa, q, inp)

		fmt.Println(result.Result)
	}
	fmt.Printf("%v elapsed\n", time.Now().Sub(start).Seconds())
	select {}
}

func decide(opa *svr.OPA, q string, inp map[string]interface{}) *svr.DecisionResult {
	ctx := context.Background()
	result, err := opa.Decision(ctx, svr.DecisionOptions{Path: q, Input: inp})
	if err != nil {
		log.Printf("obtaining decision for %q from OPA: %v", q, err)
		return nil
	}
	return result
}
