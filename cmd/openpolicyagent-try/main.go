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

	pa := rego.New( // pa: policy agent
		rego.Query("x = data.launch"),
		rego.Load([]string{"./launch.rego"}, nil),
	)

	input := jsonFileInput("./launch_input.json")
	//input := directInput()

	query := prepareEvalQuery(pa)
	rs, err := query.Eval(context.Background(), rego.EvalInput(input)) // rs: result set
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", rs[0].Bindings["x"])
	fmt.Println("Decision:", rs[0].Bindings["x"].(map[string]interface{})["decision"])
}

func jsonFileInput(inp string) interface{} {

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

func directInput() interface{} {
	return map[string]interface{}{
		"f_mass_kg":   700,
		"o2_vol_l":    1000,
		"director_ok": false,
	}
}

func prepareEvalQuery(r *rego.Rego) rego.PreparedEvalQuery {
	query, err := r.PrepareForEval(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return query
}
