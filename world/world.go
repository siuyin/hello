package world

import (
	"github.com/siuyin/hello/world/goodbye"
	"github.com/siuyin/hello/world/internal/bye"
)

func Greet() string {
	return "Hello"
}

func GoodBye() string {
	return "Goodbye"
}

func Bye2() string {
	return bye.Bye()
}

func Bye3() string {
	return goodbye.Bye()
}
