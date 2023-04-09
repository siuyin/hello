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
	input.params.FuelKg >= data.rocket.fuelLowerLimitKg
	input.params.FuelKg <  data.rocket.fuelUpperLimitKg
}

o2Go := true if {
	input.params.O2Kg >= data.rocket.o2LowerLimitKg
	input.params.O2Kg <  data.rocket.o2UpperLimitKg
}
