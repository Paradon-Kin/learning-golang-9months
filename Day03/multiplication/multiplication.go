package main

import "fmt"

func main() {
	fmt.Println("=================================================")
	fmt.Println("       Multiplication Table Generator v1.0       ")
	fmt.Println("=================================================")
	fmt.Println()

	for {
		showMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			generateSingleTable()
		case 2:
			generateRangeTable()
		case 3:
			generateCustomPattern()
		case 4:
			fmt.Println("Thank you! Goodbye")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1-4.")
		}
		fmt.Println()
	}
}

func showMenu() {
	fmt.Println("Main menu")
	fmt.Println("1. Generate Single Table")
	fmt.Println("2. Generate Range of Tables")
	fmt.Println("3. Custom Pattern")
	fmt.Println("4. Exit")
	fmt.Println()
}

func generateSingleTable() {
	var number int

	fmt.Println("Enter a number")
	fmt.Scan(&number)

	if number < 1 || number > 20 {
		fmt.Println("Number must be between 1 and 20")
		return
	}

	fmt.Printf("\nMultiplication Table of %d:\n", number)
	fmt.Println("================================")

	for i := 1; i <= 12; i++ {
		result := number * i
		fmt.Printf("%2d * %2d = %3d\n", number, i, result)
	}

	fmt.Println("================================")
}

func generateRangeTable() {
	var start, end int

	fmt.Println("Start number: ")
	fmt.Scan(&start)

	fmt.Println("End number: ")
	fmt.Scan(&end)

	if start < 1 || end > 20 || start > end {
		fmt.Println("Invalid range. Start must be <= end, both 1-29")
		return
	}

	fmt.Printf("Multiplication Table %d to %d:\n", start, end)
	fmt.Println()

	for num := start; num <= end; num++ {
		for i := 1; i <= 10; i++ {
			result := num * i
			fmt.Printf("%2d * %2d = %3d\n", num, i, result)
		}

		fmt.Println()
	}
}

func generateCustomPattern() {
	var size int

	fmt.Print("Grid size (1-12): ")
	fmt.Scan(&size)

	if size < 1 || size > 12 {
		fmt.Println("Size must be between 1 and 12")
		return
	}

	fmt.Printf("\n%d*%d Multiplication Grid:\n", size, size)
	fmt.Println()

	fmt.Print("   *")
	for i := 1; i <= size; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()
	fmt.Println(" " + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----" + "----")

	for i := 1; i <= size; i++ {
		fmt.Printf("%4d", i)

		for j := 1; j <= size; j++ {
			result := i * j
			fmt.Printf("%4d", result)
		}
		fmt.Println()
	}
	fmt.Println()
}
