package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	fmt.Println("evaluate input from rego example tests")
	ctx := context.Background()

	// Raw input data that will be used in evaluation.
	raw := `{"users": [{"id": "bob"}, {"id": "alice"}]}`
	d := json.NewDecoder(bytes.NewBufferString(raw))

	// Numeric values must be represented using json.Number.
	d.UseNumber()

	var input interface{}

	if err := d.Decode(&input); err != nil {
		panic(err)
	}

	// Create a simple query over the input.
	r := rego.New(
		rego.Query("input.users[idx].id = user_id"),
		rego.Input(input))

	//Run evaluation.
	rs, err := r.Eval(ctx)

	fmt.Printf("%#v\n%v\n%v\n", rs[0].Expressions[0].Text, len(rs), rs[1].Bindings["idx"])
	fmt.Println(err)

}
