package main

import (
	"fmt"
	"online-store/payment"
	"online-store/warehouse"
)

func main() {
	fmt.Println("==== Online Store System")

	iphone := warehouse.NewProduct("iPhone 15", 20, 29800)
	iphone.ShowDetails()

	priceToPay := iphone.GetPrice()

	success := payment.ProcessPayment(priceToPay)
	if success {
		fmt.Println("Thankyou for shopping")
	}
}
