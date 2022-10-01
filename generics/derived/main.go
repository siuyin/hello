package main

import "fmt"

type pill int

//go:generate stringer -type=pill

const (
	paracetamol pill = iota
	asprin
	cyanocobalamin
)

type drug int

//go:generate stringer -type=drug

const (
	acetaminophen drug = iota
	salicylicAsid
	vitB
)

func main() {
	fmt.Println("derived types")

	fmt.Println(code(vitB))
	fmt.Println(code(asprin))
}

func code[T ~int](item T) string {
	return fmt.Sprintf("code: %T %d %s", item, item, item)
}
