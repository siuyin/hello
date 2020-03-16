package main

import (
	"fmt"
)

type stateFn func(*int) stateFn

// sStart is the state start.
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

func main() {
	fmt.Println("Hello, playground")
	i := 0
	for sf := sStart(&i); sf != nil; sf = sf(&i) {
		//time.Sleep(time.Second)

	}
}
