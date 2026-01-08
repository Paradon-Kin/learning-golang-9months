package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	id         int
	title      string
	author     string
	isBorrowed bool
}

var bookList []Book
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("==================================")
	fmt.Println("      Mini Library System")
	fmt.Println("==================================")

	sampleBook()

	for {
		mainMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)
		fmt.Println()

		switch choice {
		case 1:
			addBook()
		case 2:
			listBooks()
		case 3:
			searchBook()
		case 4:
			borrowBook()
		case 5:
			returnBook()
		case 6:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Please Enter 1-6")
		}
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func sampleBook() {
	bookList = []Book{
		{id: 1, title: "Harry Potter and the Philosopher's Stone", author: "J.K. Rowling", isBorrowed: true},
		{id: 2, title: "1984", author: "George Orwell", isBorrowed: false},
		{id: 3, title: "To Kill a Mockingbird", author: "Harper Lee", isBorrowed: true},
		{id: 4, title: "The Great Gatsby", author: "F. Scott Fitzgerald", isBorrowed: true},
		{id: 5, title: "Pride and Prejudice", author: "Jane Austen", isBorrowed: false},
	}
}

func mainMenu() {
	fmt.Println("-------Main Menu-------")
	fmt.Println("1. Add Book")
	fmt.Println("2. List Books")
	fmt.Println("3. Search Book")
	fmt.Println("4. Borrow Book")
	fmt.Println("5. Return Book")
	fmt.Println("6. Exit")
}

func addBook() {
	var b Book

	fmt.Println("-------Add Book-------")

	if len(bookList) == 0 {
		b.id = 1
	} else {
		b.id = bookList[len(bookList)-1].id + 1
	}

	fmt.Printf("Id: %d\n", b.id)

	b.title = readLine("Title: ")
	b.author = readLine("Author: ")

	b.isBorrowed = false
	bookList = append(bookList, b)

	fmt.Printf("Book '%s' added successfully\n", b.title)
	fmt.Println()
}

func listBooks() {
	fmt.Println("-------List Book-------")

	for _, b := range bookList {
		fmt.Printf("id: %d\n", b.id)
		fmt.Printf("Title: %s\n", b.title)
		fmt.Printf("Author: %s\n", b.author)

		status := "Available"
		if b.isBorrowed {
			status = "Borrowed"
		}
		fmt.Printf("Status: %s\n", status)
	}
}

func searchBook() {
	fmt.Println("-------Search Book-------")

	searchTitle := readLine("Enter Title: ")
	searchLower := strings.ToLower(searchTitle)

	found := false

	for _, b := range bookList {
		if strings.ToLower(b.title) == searchLower {
			found = true
			fmt.Printf("ID: %d\n", b.id)
			fmt.Printf("Title: %s\n", b.title)
			fmt.Printf("Author: %s\n", b.author)

			status := "Available"
			if b.isBorrowed {
				status = "Borrowed"
			}
			fmt.Printf("Status: %s\n", status)
			fmt.Println()
		}
	}

	if !found {
		fmt.Println("Title not found")
	}
	fmt.Println()
}

func borrowBook() {
	fmt.Println("-------Borrow Book-------")

	searchTitle := readLine("Enter Title: ")

	found := false

	for index, b := range bookList {
		if strings.EqualFold(b.title, searchTitle) {
			found = true
			if b.isBorrowed {
				fmt.Println("Sorry, this book is already borrowed.")
				return
			}

			fmt.Printf("You want to borrow: %s\n", b.title)

			var choice int
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Print("Your choice: ")
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				bookList[index].isBorrowed = true
				fmt.Printf("Borrowed %s successfully\n", b.title)
				return
			default:
				fmt.Println("Operation cancelled")
				return
			}
		}
	}

	if !found {
		fmt.Println("Title not found")
	}
	fmt.Println()
}

func returnBook() {
	fmt.Println("-------Return Book-------")

	searchTitle := readLine("Enter Title: ")

	found := false

	for index, b := range bookList {
		if strings.EqualFold(b.title, searchTitle) {
			found = true

			if !b.isBorrowed {
				fmt.Println("This book is already in the library (Not borrowed).")
				return
			}

			fmt.Printf("Return Book: %s\n", b.title)

			var choice int
			fmt.Println("1. Yes")
			fmt.Println("2. Cancel")
			fmt.Print("Your choice: ")
			fmt.Scanln(&choice)

			if choice == 1 {
				bookList[index].isBorrowed = false
				fmt.Printf("Returned '%s' successfully\n", b.title)
				return
			} else {
				fmt.Println("Operation cancelled")
				return
			}
		}
	}

	if !found {
		fmt.Println("Title not found")
	}
	fmt.Println()
}
