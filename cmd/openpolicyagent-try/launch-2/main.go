package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	fmt.Println("launch-2")

	ctx := context.Background()

	// Create query that produces multiple bindings for variable.
	r := rego.New(
		rego.Query(`x = data.launch.decision`),
		rego.Package(`launch`),
		rego.Module("launch.rego", `
package launch

default fuel = "nogo"

fuel = "go" {
	input.f_mass_kg < 1000
	input.f_mass_kg >= 700
}

default oxygen = "nogo"

oxygen = "go" {
	input.o2_vol_l < 2000
	input.o2_vol_l >= 1700
}

default decision = "nogo"

decision = "go" {
	fuel == "go"
	oxygen == "go"
}

decision = "go" {
	input.director_ok
}
		`),
	)

	// Run evaluation.
	rs, err := r.Eval(ctx)

	fmt.Printf("%#v\n%v\n", rs, err)
	fmt.Println(rs[0].Bindings["x"])

}
