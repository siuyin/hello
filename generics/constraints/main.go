package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println("type constraints")

	fmt.Println(min("2b", "2a"))
	fmt.Println(min("2", "3"))
	//fmt.Println(min(3, 2.1)) // we can't compare a Float and an Integer
	fmt.Println(min(3.2, 2.1))
	fmt.Println(min(3, 2))

	a := inventory[string]{"abc", "good", 3}
	b := inventory[string]{"def", "bad", 4}
	fmt.Println(invMin(a, b))

}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type inventory[T constraints.Ordered] struct {
	sku    T
	rating T
	qty    int
}

func invMin[T constraints.Ordered](a, b inventory[T]) inventory[T] {
	if a.sku < b.sku {
		return a
	}
	return b
}
