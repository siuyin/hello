package main

import "testing"

func TestAdd2(t *testing.T) {
	if o := add2(3); o != 5 {
		t.Errorf("Unexpected value")
	}
}

func TestBalance(t *testing.T) {
	b := NewBankBalance(0)
	bal, _ := b.Deposit(10)
	if bal != 10 {
		t.Errorf("Unexpected value: %v", bal)
	}
	bal, _ = b.Withdraw(6)
	if bal != 4 {
		t.Errorf("Unexpected value: %v", bal)
	}
}
