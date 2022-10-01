package main

import "fmt"

type rect struct {
	w, h int
}
type number interface {
	int | float64
}
type areaer interface {
	area() float64
}

func main() {
	fmt.Println("Generics experiments")

	fmt.Printf("area of rectangle: %v\n", areaRect(3, 4))

	r1 := rect{3, 4}
	fmt.Printf("area of rectangle: %v\n", r1.area())

	fmt.Printf("area of circle: %v\n", areaCirc(2.0))

	c1 := circ{2.0}
	fmt.Printf("area of circle: %v\n", c1.area())

	sq1 := square{2.0}
	fmt.Printf("2x area of square: %v\n", area2x(sq1))
	fmt.Printf("2x area of circle: %v\n", area2x(c1))

	sq2 := square2[int]{2}
	fmt.Printf("generic area: %v\n", areaGeneric(sq2))

}

func areaRect(w, h int) int {
	return w * h
}
func (r rect) area() int {
	return r.w * r.h
}

const pi = 3.1415926536

func areaCirc(r float64) float64 {
	return pi * r * r
}

type circ struct {
	r float64
}

func (c circ) area() float64 {
	return pi * c.r * c.r
}

type square struct {
	s float64
}

func (s square) area() float64 {
	return s.s * s.s
}

func area2x(a areaer) float64 {
	return 2.0 * a.area()
}

type square2[T number] struct {
	r T
}

func (s square2[T]) area() T {
	return s.r * s.r
}

func areaGeneric[S square2[number], V number](s S) V {
	return s.area()
	//return 2.0
}
