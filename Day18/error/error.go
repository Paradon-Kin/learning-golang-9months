package main

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("error: Insufficient funds")

type ATM struct {
	Balance float64
}

func (a *ATM) Withdraw(amount float64) (float64, error) {
	if a.Balance >= amount {
		a.Balance -= amount
		return amount, nil
	}
	return 0, ErrorInsufficientFunds
}

func main() {
	myATM := &ATM{Balance: 600}
	fmt.Println("1. Requesting 500 THB...")
	cash, err := myATM.Withdraw(500)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Success Got %.2f THB (Remaining: %.2f)\n", cash, myATM.Balance)
	}

	fmt.Println("Requesting 500 THB...")
	cash2, err2 := myATM.Withdraw(500)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("Success got %.2f THB (Remaining %.2f)\n", cash2, myATM.Balance)
	}
}
