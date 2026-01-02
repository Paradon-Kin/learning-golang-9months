package main

import "fmt"

func main() {

	//int
	var age int = 21
	var temperature int8 = -10
	var population int64 = 7000000
	var count uint = 100

	fmt.Println("age :", age)
	fmt.Println("temperature :", temperature)
	fmt.Println("population", population)
	fmt.Println("count :", count)

	//float
	var price float32 = 99.99
	var pi float64 = 3.14159265359

	fmt.Printf("price : %.2f", price)
	fmt.Printf("pi : %.10f\n", pi)

	//string
	var name string = "Phakin"
	var lastName string = "Lee"

	fmt.Println(name)
	fmt.Println(lastName)

	//boolean
	var isStudent bool = true
	var hasGraduated bool = false

	fmt.Println("Is student :", isStudent)
	fmt.Println("Has graduated :", hasGraduated)
}
