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
