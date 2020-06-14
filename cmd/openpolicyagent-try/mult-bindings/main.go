package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	fmt.Println("multiple bindings example")

	ctx := context.Background()

	// Create query that produces multiple bindings for variable.
	r := rego.New(
		rego.Query(`a = ["ex", "am", "ple"]; x = a[_]; not p[x]`),
		rego.Package(`example`),
		rego.Module("example.rego", `package example
                p["am"] { true }
		`),
	)

	// Run evaluation.
	rs, err := r.Eval(ctx)

	fmt.Printf("%#v\n%v\n", rs, err)
	fmt.Println(rs[0].Bindings["a"])
	fmt.Println(rs[0].Bindings["x"])

}
