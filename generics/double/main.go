package main

import "fmt"

type number interface {
	int | float32 | float64
}

func main() {
	fmt.Println("generic double function")
	fmt.Println(double[int](2))
	fmt.Println(double[float64](2))
	fmt.Println(double[complex64](complex(2, 2)))
	//fmt.Println(double[inventory](inventory{"itemA", 2})) // this does not work.

	invA := inventory[int]{"itemA", 2}
	fmt.Println(invA.double()) // but this works
	invB := inventory[float32]{"itemA", 3.0}
	fmt.Println(invB.double())
	//fmt.Println(inventory{"itemA", 4.0}) // type of 4.0 should be inferred by Go but is not

	// but not this
	//invC := inventory[number]{"itemA", 4.0}

}

// I got stuck here as here is no way to overload the + operator in Go.
type inventory[T number] struct {
	sku   string
	count T
}

func (i inventory[T]) double() inventory[T] {
	j := i
	j.count = 2 * j.count
	return j
}

//func double[T int | float64 | complex64 | inventory](n T) T {
func double[T int | float64 | complex64](n T) T {
	return n + n
}
