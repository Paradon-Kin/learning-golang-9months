package main

import "fmt"

const Pi = 3.14159
const CompanyName = "Bank Plus"
const MaxStudent = 50

func main() {

	const taxRate = 0.07
	const vat = 0.07

	fmt.Println("Pi:", Pi)
	fmt.Println("Company:", CompanyName)
	fmt.Println("Tax Rate:", taxRate)

	//Using constants in calculations
	var price float64 = 1000.0

	totalWithTax := price + (price * taxRate)
	fmt.Printf("Price: %.2f\n", price)
	fmt.Printf("Total with Tax: %.2f\n", totalWithTax)

	//constants vs variable
	//var discount float64 = 0.10
	const discount float64 = 0.10

	//discount = 0.20

	//fmt.Printf("Discount: %.2f\n", discount)
}
