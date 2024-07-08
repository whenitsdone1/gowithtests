package Wallet

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.viewBalance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	t.Run("Testing Withdraw...", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(8)
		if err != nil {
			t.Fatal("Panic!")
		}
		got := wallet.viewBalance()
		want := Bitcoin(2)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
	t.Run("Testing too few funds...", func(t *testing.T) {
		wallet := Wallet{2}
		got := wallet.Withdraw(8)
		if !errors.Is(got, lackFundsError) {
			t.Errorf("got %s want %s", got, lackFundsError)
		}

	})

}
