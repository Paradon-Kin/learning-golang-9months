package main

import "fmt"

func main() {
	fmt.Println("====Classic for looop====")
	fmt.Println()

	// //ex1
	// fmt.Println("Count from 1 to 5")
	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("%d\n", i)
	// }
	// println()

	// //ex2 count backwards
	// for i := 10; i >= 0; i-- {
	// 	fmt.Printf("%d\n", i)
	// }
	// println()

	// //ex3 skip by 2
	// for i := 0; i <= 10; i += 2 {
	// 	fmt.Printf("%d\n", i)
	// }
	// println()

	// //ex4 multiplication table
	// for i := 1; i <= 10; i++ {
	// 	result := 7 * i
	// 	fmt.Printf("7 * %d = %d\n", i, result)
	// }
	// fmt.Println()

	// //while-style loop
	// count := 1
	// for count <= 5 {
	// 	fmt.Printf("Count: %d\n", count)
	// 	count++
	// }
	// println()

	//ex keep asking until valid input
	// var number int
	// fmt.Println("Guess a number between 1 and 10:")

	// for {
	// 	fmt.Print("Your guess: ")
	// 	fmt.Scan(&number)

	// 	if number >= 1 && number <= 10 {
	// 		fmt.Printf("valid you entered: %d\n", number)
	// 		break
	// 	}
	// 	fmt.Println("Invalid! try again.")
	// }

	//ex stop when found
	// fmt.Println("fine first numer divisible by 7:")
	// for i := 1; i <= 100; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Printf("found: %d\n", i)
	// 		break
	// 	}
	// }

	//ex exit on condition
	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += 1
	// 	fmt.Printf("i=%d, sum=%d\n", i, sum)

	// 	if sum > 100 {
	// 		fmt.Println("Sum exceeded 100! Stopping.")
	// 		break
	// 	}
	// }

	//skip odd numbers
	// for i := 0; i <= 20; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}

	// 	fmt.Printf("%d", i)
	// }
	// fmt.Println()

	//multiplication table
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			result := i * j
			fmt.Printf("%4d", result)
		}
		fmt.Println()
	}
	fmt.Println()

	//pattern printing
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
