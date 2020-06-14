package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	fmt.Println("simple binding taken from rego example test")
	r := rego.New(rego.Query("x = 2"))
	rs, err := r.Eval(context.Background())
	fmt.Println(rs[0].Bindings["x"])
	fmt.Println(err)
}
