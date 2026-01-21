package main

import "fmt"

type Book struct {
	title       string
	author      string
	isAvailable bool
}

type Library struct {
	name    string
	shelves []*Book
}

func NewBook(t string, a string) *Book {
	return &Book{
		title:       t,
		author:      a,
		isAvailable: true,
	}
}

func (b *Book) Borrow() {
	if b.isAvailable {
		b.isAvailable = false
		fmt.Printf("Borrowed book Title:%s successfully\n", b.title)
	} else {
		fmt.Println("This book is borrowed")
	}
}

func (b *Book) Return() {
	b.isAvailable = true
	fmt.Printf("Return book Title:%s successfully\n", b.title)
}

func (b *Book) GetStatus() bool {
	return b.isAvailable
}

func (l *Library) AddBook(b *Book) {
	l.shelves = append(l.shelves, b)
	fmt.Printf("[Library] Added '%s' to shelves.\n", b.title)
}

func (l *Library) FindAndBorrow(bookName string) {
	found := false
	for i := range l.shelves {
		if l.shelves[i].title == bookName {
			found = true
			l.shelves[i].Borrow()
			break
		}
	}
	if !found {
		fmt.Println("Book not found")
	}
}

func (l *Library) ShowAllBook() {
	for _, i := range l.shelves {
		if i.isAvailable {
			fmt.Printf("[Title:%s |Author:%s |Status: Available]\n", i.title, i.author)
		} else {
			fmt.Printf("[Title:%s |Author:%s |Status: Borrowed]\n", i.title, i.author)
		}
	}
	fmt.Println("--------------------------------------------------")
}

func main() {
	lib := Library{name: "IT Library"}

	book1 := NewBook("Romeo and Juliet", "William Shakespeare")
	book2 := NewBook("Emma", "Jane Austen")
	book3 := NewBook("A Tale of Two Cities", "Charles Dickens")

	lib.AddBook(book1)
	lib.AddBook(book2)
	lib.AddBook(book3)

	lib.FindAndBorrow("Romeo and Juliet")
	lib.FindAndBorrow("Romeo and Juliet")
	lib.FindAndBorrow("Emma")

	lib.ShowAllBook()

}
