package payment

import "fmt"

func ProcessPayment(amount float64) bool {
	if amount <= 0 {
		fmt.Println("Payment Error: Invalid amount")
	}
	fmt.Printf("[Payment] Processing payment: %.2f THB... Success\n", amount)
	return true
}
