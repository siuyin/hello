package main

import (
	"fmt"
)

type stateFn func(*int) stateFn

// sStart is the start state.
// i can be thought as the environment surrounding the state machine.
// In this example, i is used for terminating the state machine.
// State machines in embedded electronics typically do not terminate.
func sStart(i *int) stateFn {
	fmt.Println("start")
	return sOpen(i)
}
func sOpen(i *int) stateFn {
	fmt.Println("open")
	return sClose(i)
}
func sClose(i *int) stateFn {
	*i++
	if *i > 10 {
		return nil
	}
	fmt.Printf("close %d\n", *i)
	return sOpen(i)
}
func runStateMachine(i *int) {
	for sf := sStart(i); sf != nil; sf = sf(i) {
	}
}

func main() {
	fmt.Printf("State machine demo\n\n")
	i := 0
	runStateMachine(&i)
	fmt.Println("\nDemo completed")
}
