package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	fmt.Println("open policy agent as a go library")
	r := rego.New(
		rego.Query("x = data.launch"),
		rego.Load([]string{"./launch.rego"}, nil))

	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadFile("./launch_input.json")
	if err != nil {
		log.Fatal(err)
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		log.Fatal(err)
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", rs[0].Bindings["x"])
	fmt.Println("Decision:", rs[0].Bindings["x"].(map[string]interface{})["decision"])

}
