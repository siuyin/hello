package main

import "fmt"

func main() {
	fmt.Println("Pure functions and state")
}

// Below are pure functions. They can be put into a package.
func add2(i int) int {
	return i + 2
}
func add(a, b int) int {
	return a + b
}

// BankBalance is a stateful type. It can make use of pure functions
// from the package above.
type BankBalance struct {
	Balance int
}

func NewBankBalance(i int) *BankBalance {
	return &BankBalance{i}
}
func (b *BankBalance) Deposit(i int) (int, error) {
	b.Balance = add(b.Balance, i) // roundabout way of adding a number
	return b.Balance, nil
}
func (b *BankBalance) Withdraw(i int) (int, error) {
	if i > b.Balance {
		return b.Balance, fmt.Errorf("Insufficient funds, balance: %v, trying to withdraw %v", b.Balance, i)
	}
	b.Balance = add(b.Balance, -i)
	return b.Balance, nil
}
