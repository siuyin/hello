package world

import "testing"

func TestGreet(t *testing.T) {
	if g := Greet(); g != "Hello" {
		t.Errorf("Unexpected value: %s, expected Hello", g)
	}
}

func TestGoodBye(t *testing.T) {
	if b := GoodBye(); b != "Goodbye" {
		t.Errorf("Unexpected value: %s, expected Goodbye", b)
	}
}
