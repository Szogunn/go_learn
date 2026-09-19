package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p\n", &wallet.balance)

	want := Bitcoin(10)

	if got != want {
		t.Errorf("want %d got %d", got, want)
	}
}

func TestBitcoin(t *testing.T) {
	t.Run("Bitcoin String", func(t *testing.T) {
		btc := Bitcoin(10)
		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
