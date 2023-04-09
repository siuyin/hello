package rocket
import future.keywords

default launch := false
launch := true if {
	fuelGo
	o2Go
	input.params.AvionicsGo
	not input.params.FlightDirectorNoGo
}

fuelGo := true if {
	input.params.FuelKg >= 100
	input.params.FuelKg <  120
}

o2Go := true if {
	input.params.O2Kg >= 200
	input.params.O2Kg < 220
}
