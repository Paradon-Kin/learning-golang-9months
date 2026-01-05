package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}

	for index, fruit := range fruits {
		fmt.Printf("%d: %s\n", index, fruit)
	}
	fmt.Println()

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
	fmt.Println()

	for index := range fruits {
		fmt.Printf("Index: %d\n", index)
	}
	fmt.Println()

	numbers := []int{10, 20, 30, 40, 50}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Sum of %v = %d\n", numbers, sum)

}
