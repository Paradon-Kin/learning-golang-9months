package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type product struct {
	name  string
	price float64
	stock int
}

var productList []product

func main() {

	sampleProduct()

	for {
		mainMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		scanner.Scan()

		switch choice {
		case 1:
			addProduct()
		case 2:
			updateStock()
		case 3:
			viewProduct()
		default:
			fmt.Println("Please enter 1-4")
		}
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func sampleProduct() {
	productList = []product{
		{name: "iPhone 17", price: 29800, stock: 38},
		{name: "iPhone 17 pro", price: 39800, stock: 48},
		{name: "iPhone 17 proMax", price: 45800, stock: 56},
	}
}

func mainMenu() {
	fmt.Println("1. Add product")
	fmt.Println("2. Update stock")
	fmt.Println("3. View product")
	fmt.Println("4. Exit")
}

func addProduct() {
	var p product

	fmt.Println("----- Add product -----")
	p.name = readLine("Name: ")

	fmt.Print("Price: ")
	fmt.Scan(&p.price)

	fmt.Print("Stock: ")
	fmt.Scan(&p.stock)
	scanner.Scan()

	if p.price < 1 || p.stock < 1 {
		fmt.Println("Invalid")
	}

	productList = append(productList, p)

	fmt.Printf("Add %s price %.1f(s)\n", p.name, p.price)

}

func updateStock() {
	productName := readLine("Name: ")

	found := false

	for i, p := range productList {
		if strings.EqualFold(p.name, productName) {
			found = true
			fmt.Printf("Product %s: Price %.1f: Stock: %d\n", p.name, p.price, p.stock)

			var Newstock int
			fmt.Print("Update: ")
			fmt.Scan(&Newstock)
			scanner.Scan()
			updatefunc(&productList[i], Newstock)
			return
		}
	}

	if !found {
		fmt.Println("Not found product")
	}
}

func updatefunc(p *product, qty int) {
	p.stock = p.stock + qty
	fmt.Println("Update stock successfully")
}

func viewProduct() {
	for _, i := range productList {
		fmt.Printf("Product: %s : Price: %.1f : Stock: %d\n", i.name, i.price, i.stock)
	}
}
