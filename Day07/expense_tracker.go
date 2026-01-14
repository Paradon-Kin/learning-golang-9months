package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Transaction struct {
	id          int
	description string
	amount      float64
	class       string
}

var transaction []Transaction
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("======== Expense Tracker ========")

	sample()

	for {
		mainMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEntry()
		case 2:
			viewAll()
		case 3:
			cal()
		default:
			fmt.Println("please enter 1-5")
		}

	}
}

func sample() {
	transaction = []Transaction{
		{id: 1, description: "kao man kai", amount: 40, class: "Expense"},
		{id: 2, description: "Kao grapao", amount: 45, class: "Expense"},
	}
}

func mainMenu() {
	fmt.Println("-------- Main Menu --------")
	fmt.Println("1. Add Entry")
	fmt.Println("2. View all")
	fmt.Println("3. View Balance")
	fmt.Println("4. Delete Entry")
	fmt.Println("5. Exit")
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func addEntry() {
	var t Transaction

	fmt.Println("-------- Add Entry --------")

	if len(transaction) == 0 {
		t.id = 1
	} else {
		t.id = transaction[len(transaction)-1].id + 1
	}

	fmt.Printf("ID: %d\n", t.id)

	t.description = readLine("Description: ")

	fmt.Print("Amount: ")
	fmt.Scanln(&t.amount)
	if t.amount < 1 {
		fmt.Println("Invalid")
		return
	}

	t.class = readLine("Class: ")

	transaction = append(transaction, t)

	fmt.Printf("Add %s successfully\n", t.description)
	fmt.Println()
}

func viewAll() {
	fmt.Println("------- View All --------")

	for _, t := range transaction {
		fmt.Printf("ID: %d\n", t.id)
		fmt.Printf("Description: %s\n", t.description)
		fmt.Printf("Amount: %.1f\n", t.amount)
		fmt.Printf("Class: %s\n", t.class)
	}
}

func cal() {
	fmt.Println("-------- View Balance --------")
	var income float64
	for index, t := range transaction {
		if transaction[index].class == "expense" {
			income += t.amount
		}
		fmt.Printf("sum %.1f\n", income)
	}
}
