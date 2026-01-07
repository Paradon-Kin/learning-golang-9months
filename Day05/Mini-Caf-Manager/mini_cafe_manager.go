package main

import "fmt"

type Drink struct {
	id       int
	name     string
	price    int
	category string
}

var menulist []Drink

func main() {
	fmt.Println("======================================")
	fmt.Println("          Mini Cafe Manager           ")
	fmt.Println("======================================")
	fmt.Println()

	sampleMenu()

	for {
		mainMenu()

		var choice int
		fmt.Print("Your Choice (1-6): ")
		fmt.Scan(&choice)
		fmt.Println()

		switch choice {
		case 1:
			showMenu()
		case 2:
			addDrink()
		case 3:
			editPrice()
		case 4:
			deleteDrink()
		case 5:
			cashier()
		case 6:
			fmt.Println("Goodbye! Closing system...")
			return
		default:
			fmt.Println("Please enter your choice between 1 - 6")
		}
	}
}

func sampleMenu() {
	menulist = []Drink{
		{id: 1, name: "Latte", price: 130, category: "Hot"},
		{id: 2, name: "Espresso", price: 120, category: "Hot"},
		{id: 3, name: "Americano", price: 150, category: "Iced"},
	}
}

func mainMenu() {
	fmt.Println("-------------Main Menu-----------")
	fmt.Println("1. Show Menu")
	fmt.Println("2. Add Drink")
	fmt.Println("3. Edit Price")
	fmt.Println("4. Delete Drink")
	fmt.Println("5. Cashier")
	fmt.Println("6. Exit")
	fmt.Println("----------------------------------")
	fmt.Println()
}

func showMenu() {
	fmt.Println("---------- All Menu ----------")

	for _, d := range menulist {
		fmt.Printf("#%d Name: %s\n", d.id, d.name)
		fmt.Printf("Price: %d\n", d.price)
		fmt.Printf("Category: %s\n", d.category)
		fmt.Println("------------------------------")
	}
	fmt.Println()
}

func addDrink() {
	var d Drink

	fmt.Println("----------- Add Drink --------------")

	if len(menulist) == 0 {
		d.id = 1
	} else {
		d.id = menulist[len(menulist)-1].id + 1
	}

	fmt.Printf("Generated ID: %d\n", d.id)

	fmt.Print("Drink name: ")
	fmt.Scan(&d.name)

	fmt.Print("Price: ")
	fmt.Scan(&d.price)
	if d.price < 1 {
		fmt.Println("Invalid price")
		return
	}

	fmt.Println("Category")
	fmt.Println("1. Hot")
	fmt.Println("2. Iced")
	fmt.Println("3. Frappe")

	var category int
	fmt.Print("Your choice: ")
	fmt.Scan(&category)

	switch category {
	case 1:
		d.category = "Hot"
	case 2:
		d.category = "Iced"
	case 3:
		d.category = "Frappe"
	default:
		fmt.Println("Please enter 1 - 3")
		return
	}

	menulist = append(menulist, d)

	fmt.Printf("Add menu '%s' successfully\n", d.name)
	fmt.Println("---------------------------------")
	fmt.Println()
}

func editPrice() {
	fmt.Println("---- Edit Price ----")

	fmt.Println("Enter the menu name to edit price")

	var searchMenu string
	var newPrice int

	fmt.Print("Menu name: ")
	fmt.Scan(&searchMenu)

	found := false

	for i := range menulist {
		if menulist[i].name == searchMenu {
			found = true
			fmt.Printf("Current price = %d\n", menulist[i].price)

			fmt.Print("New price: ")
			fmt.Scan(&newPrice)

			menulist[i].price = newPrice
			fmt.Printf("Price for '%s' updated successfully\n", menulist[i].name)
			break
		}
	}
	fmt.Println()

	if !found {
		fmt.Println("Menu not found")
	}
}

func deleteDrink() {
	var searchName string

	fmt.Println("--------- Delete Drink ---------")

	fmt.Println("Enter the menu name to delete")
	fmt.Print("Menu name: ")
	fmt.Scan(&searchName)

	found := false

	for index, d := range menulist {
		if d.name == searchName {
			found = true

			var choice int
			fmt.Printf("Found '%s'. Are you sure you want to delete?\n", d.name)
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Print("Your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				menulist = append(menulist[:index], menulist[index+1:]...)
				fmt.Printf("Menu '%s' deleted successfully\n", d.name)
				return
			case 2:
				fmt.Println("Deletion cancelled")
				return
			}
		}
	}

	if !found {
		fmt.Println("Menu not found")
	}

	fmt.Println()
}

func cashier() {
	fmt.Println("------- Cashier -------")

	var menuName string
	fmt.Print("Enter menu name: ")
	fmt.Scan(&menuName)

	found := false

	for m := range menulist {
		if menulist[m].name == menuName {
			found = true

			var quantity int
			fmt.Print("Quantity: ")
			fmt.Scan(&quantity)
			if quantity <= 0 {
				fmt.Println("Invalid quantity")
				return
			}

			price := menulist[m].price
			total := price * quantity

			fmt.Printf("Total amount: %d\n", total)

			var money int
			fmt.Print("Cash received: ")
			fmt.Scan(&money)

			if money < total {
				fmt.Println("Not enough money")
				fmt.Printf("Missing: %d\n", total-money)
				return
			}

			change := money - total

			fmt.Printf("Change: %d\n", change)
			fmt.Println("Thank you!")
			break
		}
	}

	if !found {
		fmt.Println("Menu not found")
		return
	}
	fmt.Println()
}
