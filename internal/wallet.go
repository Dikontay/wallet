package internal

import (
	"errors"
	"sync"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
	mu      sync.Mutex
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient balance")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance
}
