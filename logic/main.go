package main

import "fmt"

func main() {
	fmt.Println("propositional logic examples")
	// logic symbols representing propositions which are either true or false.
	var aa, ab, ac, ad, ba, bb, bc, bd, ca, cb, cc, cd, da, db, dc, dd bool
	aa = true
	bd = true
	cb = true
	dc = true
	println(aa, ab, ac, ad, ba, bb, bc, bd, ca, cb, cc, cd, da, db, dc, dd)

	kb := !bc &&
		(!bb || dd) &&
		(!ac || da) &&
		(dd || !ba) &&
		(!ca || bb) &&
		((bb || cc) || dc)

	if kb &&
		(aa && bd && cb && dc) {
		println("verified")
	}

}
