package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	money Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.money += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// nullのようにnilである値にアクセスしようとすると、ランタイムパニックがスローされます。
// Withdrawの戻り値の型はインターフェイスであるerrorになるため、エラーは nilになる可能性があります。
// 引数を取る関数、またはインターフェイスである値を返す関数がある場合、それらは nillable(潰しが利かない) になる可能性があります。
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.money-amount < 0 {
		return ErrInsufficientFunds
	}
	w.money -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.money
}
