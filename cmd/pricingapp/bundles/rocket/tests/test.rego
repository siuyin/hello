package rocket
import future.keywords

test_fuelGo_good  if {
	fuelGo with input as {"params":{"FuelKg": data.rocket.fuelLowerLimitKg+0.1}}
}
test_fuelGo_low if {
	not fuelGo with input as {"params":{"FuelKg": data.rocket.fuelLowerLimitKg-0.1}}
}
test_fuelGo_excessive if {
	not fuelGo with input as {"params":{"FuelKg": data.rocket.fuelUpperLimitKg+0.1}}
}

test_o2Go_good if {
	o2Go with input as {"params":{"O2Kg": data.rocket.o2LowerLimitKg+0.1}}
}
test_o2Go_low if {
	not o2Go with input as {"params":{"O2Kg": data.rocket.o2LowerLimitKg - 0.1}}
}
test_o2Go_excessive if {
	not o2Go with input as {"params":{"O2Kg": data.rocket.o2UpperLimitKg + 0.1}}
}

test_launch_go if {
	launch with fuelGo as true
		with o2Go as true
		with input.params.AvionicsGo as true
		with input.params.FlightDirectorNoGo as false
}
test_launch_nogo_flight_override if {
	not launch with fuelGo as true
		with o2Go as true
		with input.params.AvionicsGo as true
		with input.params.FlightDirectorNoGo as true
}
test_launch_nogo_with_one_nogo_fuel {
	not launch with fuelGo as false 
		with o2Go as true
		with input.params.AvionicsGo as true
}
test_launch_nogo_with_one_nogo_avionics {
	not launch with fuelGo as true
		with o2Go as true
		with input.params.AvionicsGo as false
}
test_launch_nogo_with_one_nogo_and_flight_override {
	not launch with fuelGo as true
		with o2Go as true
		with input.params.AvionicsGo as false
		with input.params.FlightDirectorNoGo as false
}
test_launch_nogo_with_one_O2Kg_undefined {
	not launch with fuelGo as true
		with input.params.AvionicsGo as true
}
test_launch_nogo_with_all_undefined_and_flight_go {
	not launch with input.params.FlightDirectorNoGo as false
}


