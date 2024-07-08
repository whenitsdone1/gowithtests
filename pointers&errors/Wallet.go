package Wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	Balance Bitcoin
}

var lackFundsError = errors.New("Lack Funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) viewBalance() Bitcoin {
	return w.Balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance {
		return lackFundsError
	}
	w.Balance -= amount
	return nil
}
