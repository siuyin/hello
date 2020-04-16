package main

import (
	"fmt"
	"testing"

	"github.com/siuyin/hello/world"
)

func ExampleHello() {
	fmt.Println(world.Greet())
	// Output:
	// Hello
}

func TestAdd(t *testing.T) {
	if 1+1 != 2 {
		t.Error("failed test")
	}
}

func TestSub(t *testing.T) {
	if 3-1 != 2 {
		t.Error("failed test")
	}
}
