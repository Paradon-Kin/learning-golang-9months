package main

import "errors"

func Withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return 0, errors.New("invalid amount")
	}

	if balance < amount {
		return 0, errors.New("insufficient funds")
	}

	balance -= amount
	return balance, nil
}
