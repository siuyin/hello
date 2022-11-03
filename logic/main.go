package main

import "fmt"

func main() {
	fmt.Println("propositional logic examples")
	// Logic symbols representing propositions which are either true or false.
	// Below is one example or "model" of the logic problem.
	var aa, ab, ac, ad, ba, bb, bc, bd, ca, cb, cc, cd, da, db, dc, dd bool
	aa = true
	bd = true
	cb = true
	dc = true
	println(aa, ab, ac, ad, ba, bb, bc, bd, ca, cb, cc, cd, da, db, dc, dd)

	knowledge := !bc &&
		(!bb || dd) &&
		(!ac || da) &&
		(dd || !ba) &&
		(!ca || bb) &&
		((bb || cc) || dc)

	// if knowlege is true when evaluated against a model AND our query also evaluates to true, then our query entails the knowledge.
	if knowledge &&
		(aa && bd && cb && dc) {
		println("verified")
	}

}
