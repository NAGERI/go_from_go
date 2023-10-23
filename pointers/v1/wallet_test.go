package main

import (
	"testing"
)
/**
 *  We use structs to manage state, exposing methods to let users change the state in a way that you can control */
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	
	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d instead of %d", got, want )
	}
}