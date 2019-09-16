package main

import (
	"fmt"

	"github.com/siuyin/hello/world"
)

func main() {
	fmt.Println(world.Greet(), "world!")
	fmt.Println(world.GoodBye(), "world!")
}
