package internal

import (
	"sync"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient balance")
		assertBalance(t, wallet, startingBalance)
	})

	// Additional tests to cover concurrent access
	t.Run("Concurrent Deposits", func(t *testing.T) {
		wallet := Wallet{}
		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			wallet.Deposit(Bitcoin(10))
			wg.Done()
		}()
		go func() {
			wallet.Deposit(Bitcoin(20))
			wg.Done()
		}()
		wg.Wait()

		assertBalance(t, wallet, Bitcoin(30))
	})
	// Additional tests can be written for concurrent withdrawals and a mix of operations
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got.Error(), want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}
