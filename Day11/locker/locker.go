package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type locker struct {
	id      int
	size    string
	isTaken bool
	owner   string
}

var scanner = bufio.NewScanner(os.Stdin)

var lockers []locker

func main() {
	sampleData()

	for {
		mainMenu()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
		scanner.Scan() // Clear buffer
		fmt.Println()

		switch choice {
		case 1:
			viewLockers()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			fmt.Println("Thank you! Goodbye.")
			return
		default:
			fmt.Println("Please enter a number between 1-4")
		}
	}
}

func sampleData() {
	lockers = []locker{
		{id: 1, size: "L", isTaken: true, owner: "Jack"},
		{id: 2, size: "L", isTaken: false, owner: ""},
		{id: 3, size: "M", isTaken: true, owner: "Sindy"},
		{id: 4, size: "S", isTaken: false, owner: ""},
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu() {
	fmt.Println("\n==== Main Menu ====")
	fmt.Println("1. View Lockers")
	fmt.Println("2. Deposit Item")
	fmt.Println("3. Withdraw Item")
	fmt.Println("4. Exit")
}

func viewLockers() {
	fmt.Println("--- Locker Status ---")
	for _, l := range lockers {
		var status string
		if l.isTaken {
			status = "Taken by " + l.owner
		} else {
			status = "Empty"
		}

		fmt.Printf("ID: %d | Size: %s | %s\n", l.id, l.size, status)
	}
	fmt.Println("---------------------")
}

func deposit() {
	var inputID int
	fmt.Print("Enter Locker ID: ")
	fmt.Scan(&inputID)
	scanner.Scan()

	found := false // 🚩 1. ตัวแปรเช็คว่าเจอ ID ไหม

	for i := range lockers {
		if lockers[i].id == inputID {
			found = true // เจอแล้ว!

			if lockers[i].isTaken {
				// แจ้งเตือนว่าไม่ว่าง และบอกด้วยว่าเป็นของใคร
				fmt.Printf("Error: Locker %d is already taken by %s.\n", lockers[i].id, lockers[i].owner)
			} else {
				owner := readLine("Enter Owner Name: ")
				dep(&lockers[i], owner) // ส่ง Pointer ไป
				fmt.Println("✅ Deposit successful!")
			}
			break // 🚀 2. เจอแล้วหยุดหาทันที (Optimization)
		}
	}

	// 3. ถ้าหาจนจบแล้วไม่เจอ found
	if !found {
		fmt.Println("Error: Locker ID not found.")
	}
}

func dep(l *locker, o string) {
	l.isTaken = true
	l.owner = o
}

func withdraw() {
	var inputID int
	fmt.Print("Enter Locker ID: ")
	fmt.Scan(&inputID)
	scanner.Scan()

	found := false

	for i := range lockers {
		if lockers[i].id == inputID {
			found = true

			if lockers[i].isTaken {
				w(&lockers[i]) // ส่ง Pointer ไป
				fmt.Println("✅ Withdraw successful! Locker is now empty.")
			} else {
				fmt.Println("Error: This locker is already empty.")
			}
			break // เจอแล้วหยุดหา
		}
	}

	if !found {
		fmt.Println("Error: Locker ID not found.")
	}
}

func w(l *locker) {
	l.isTaken = false
	l.owner = ""
}
