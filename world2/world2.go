package world2

import (
	"github.com/siuyin/hello/world"
	"github.com/siuyin/hello/world/goodbye"
	"github.com/siuyin/hello/world/internal/bye"
)

func Bye() string {
	return goodbye.Bye()
}

func Bye2() string {
	return world.Bye2()
}

// this will not compile. Error message: use of internal package github.com/siuyin/hello/world/internal/bye not allowed
func Bye3() string {
	return bye.Bye()
}
