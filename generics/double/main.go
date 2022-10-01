package main

import "fmt"

func main() {
	fmt.Println("generic double function")
	fmt.Println(double[int](2))
	fmt.Println(double[float64](2))
	fmt.Println(double[complex64](complex(2, 2)))
	//fmt.Println(double[inventory](inventory{"itemA", 2})) // this does not work.

	invA := inventory[int]{"itemA", 2}
	fmt.Println(invA.double()) // but this works

}

// I got stuck here as here is no way to overload the + operator in Go.
type inventory[T int | float64] struct {
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
