package main

import (
	"fmt"

	"github.com/siuyin/hello/world"
)

func main() {
	fmt.Println(world.Greet(), "World!")
	var n int
	fmt.Println(n)
	fmt.Println(sub(3, 2))
	fmt.Println(world.GoodBye(), "world!")
}

func sub(a, b int) int {
	//fmt.Println("Sub")
	return a - b
}
