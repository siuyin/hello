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
		rego.Load([]string{"./launch.rego"}, nil),
	)

	input := loadInput("./launch_input.json")

	query := prepareEvalQuery(r)
	rs, err := query.Eval(context.Background(), rego.EvalInput(input)) // rs: result set
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", rs[0].Bindings["x"])
	fmt.Println("Decision:", rs[0].Bindings["x"].(map[string]interface{})["decision"])
}

func loadInput(inp string) interface{} {

	bs, err := ioutil.ReadFile(inp)
	if err != nil {
		log.Fatal(err)
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		log.Fatal(err)
	}
	return input
}

func prepareEvalQuery(r *rego.Rego) rego.PreparedEvalQuery {
	query, err := r.PrepareForEval(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return query
}
