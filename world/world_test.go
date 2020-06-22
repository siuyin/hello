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

func TestBye2(t *testing.T) {
	if s := Bye2(); s != "Goodbye" {
		t.Errorf("Unexpected value: %s, expected Goodbye", s)
	}
}

func TestBye3(t *testing.T) {
	if s := Bye3(); s != "Goodbye" {
		t.Errorf("Unexpected value: %s, expected Goodbye", s)
	}
}
