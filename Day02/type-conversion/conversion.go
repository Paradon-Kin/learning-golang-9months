package main

import "fmt"

func main() {

	//int to float
	var x int = 10
	var y float64

	y = float64(x)

	fmt.Printf("x (int): %d\n", x)
	fmt.Printf("y (float64): %.2f\n", y)

	//float to int
	var price float64 = 99.99
	var roundedPrice int

	roundedPrice = int(price)
	fmt.Printf("Original price: %.2f\n", price)
	fmt.Printf("Rounded price: %d\n", roundedPrice)

	//practical example
	var totalScore int = 267
	var subject int = 5

	wrongAverage := totalScore / subject
	fmt.Printf("Wrong average: %d\n", wrongAverage)

	correctAverage := float64(totalScore) / float64(subject)
	fmt.Printf("Correct average: %.2f\n", correctAverage)

}
