package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsuficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsuficientFunds = errors.New("cannot withdraw, insuficient funds")
