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

	a := inventory[string, uint]{"abc", "good", 3}
	b := inventory[string, uint]{"def", "bad", 4}
	fmt.Println(invSKUMin(a, b))
	fmt.Println(a.qtyMax(b))

}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type inventory[T constraints.Ordered, N constraints.Unsigned] struct {
	sku    T
	rating T
	qty    N
}

func invSKUMin[T constraints.Ordered, N constraints.Unsigned](a, b inventory[T, N]) inventory[T, N] {
	if a.sku < b.sku {
		return a
	}
	return b
}

func (i inventory[T, N]) qtyMax(b inventory[T, N]) inventory[T, N] {
	if i.qty < b.qty {
		return b
	}
	return i
}

func (i inventory[T, N]) String() string {
	return fmt.Sprintf("sku: %v, rating: %v, qty: %v", i.sku, i.rating, i.qty)
}
