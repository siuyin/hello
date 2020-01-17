package main

import (
	"fmt"

	"github.com/siuyin/hello/world"
)

func main() {
	fmt.Println(world.Greet(), "world!")
	var n int
	fmt.Println(n)
	fmt.Println(world.GoodBye(), "world!")
}
