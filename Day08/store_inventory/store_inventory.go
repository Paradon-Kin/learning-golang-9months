package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	stock := make(map[string]int)

	stock["iPhone 16"] = 20
	stock["iPhone 17"] = 25
	stock["iPhone 18"] = 30
	stock["iPhone 19"] = 10

	for {
		mainMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		scanner.Scan()

		switch choice {
		case 1:
			name := readLine("Product Name: ")

			var amount int
			fmt.Print("Amount: ")
			fmt.Scan(&amount)
			scanner.Scan()

			stock[name] += amount
			fmt.Printf("Added %d %s(s)\n", amount, name)
		case 2:
			name := readLine("Product name to sell: ")
			qty, exists := stock[name]
			if !exists {
				fmt.Println("Error: Product not found")
				break
			}

			var amount int
			fmt.Print("Amount to sell: ")
			fmt.Scan(&amount)
			scanner.Scan()

			if amount > qty {
				fmt.Println("Error: Not enough stock! (Available:", qty, ")")
			} else {
				stock[name] -= amount
				fmt.Printf("Sold %d %s(s). Remaining: %d\n", amount, name, stock[name])
				if stock[name] == 0 {
					delete(stock, name)
					fmt.Println("Stock empty, removed from list.")
				}
			}
		case 3:
			if len(stock) == 0 {
				fmt.Println("Empty stock")
			} else {
				for name, qty := range stock {
					fmt.Printf("- %s : %d\n", name, qty)
				}
			}
		case 4:
			fmt.Println("Byeee")
			return
		default:
			fmt.Println("Please enter 1-4")
		}
	}

}

func mainMenu() {
	fmt.Println("1. Add Item")
	fmt.Println("2. Sell Item")
	fmt.Println("3. Check Stock")
	fmt.Println("4. Exit")
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
