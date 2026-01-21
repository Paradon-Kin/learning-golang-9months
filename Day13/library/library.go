package main

import (
	"fmt"
)

type Book struct {
	Title       string
	Author      string
	isAvailable bool // Private: ห้ามแก้ตรงๆ
}

type Library struct {
	Name    string
	Shelves []*Book // เก็บ Pointer ของหนังสือ
}

// Constructor สร้างหนังสือ
func NewBook(title string, author string) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		isAvailable: true,
	}
}

// Method ยืมหนังสือ
func (b *Book) Borrow() {
	if !b.isAvailable {
		fmt.Printf("❌ Error: '%s' is already borrowed.\n", b.Title)
	} else {
		b.isAvailable = false
		fmt.Printf("✅ Success: You borrowed '%s'.\n", b.Title)
	}
}

// Method คืนหนังสือ
func (b *Book) Return() {
	b.isAvailable = true
}

func (b *Book) GetStatus() bool {
	return b.isAvailable
}

// 🛠️ แก้ไข 1: รับ b เข้ามาแล้วใส่เข้าชั้นวางเลย ไม่ต้องสร้างใหม่
func (l *Library) AddBook(b *Book) {
	l.Shelves = append(l.Shelves, b)
	fmt.Printf("[Library] Added '%s' to shelves.\n", b.Title)
}

// 🛠️ แก้ไข 2: เรียก method Borrow จากตัวหนังสือในชั้นวาง
func (l *Library) FindAndBorrow(bookName string) {
	found := false
	for i := range l.Shelves {
		// เช็คชื่อหนังสือ
		if bookName == l.Shelves[i].Title {
			found = true
			l.Shelves[i].Borrow() // เรียก Method ของหนังสือเล่มนั้น
			break                 // เจอแล้วหยุดหา
		}
	}

	if !found {
		fmt.Printf("🔍 Not found: Book '%s' doesn't exist.\n", bookName)
	}
}

// เพิ่มฟังก์ชันแสดงรายการหนังสือทั้งหมด (ตามโจทย์)
func (l *Library) ShowAllBooks() {
	fmt.Println("\n--- 📚 Library Catalog ---")
	for _, b := range l.Shelves {
		status := "Available"
		if !b.GetStatus() { // ใช้ Getter เพราะ isAvailable เป็น private
			status = "Borrowed"
		}
		fmt.Printf("- %s (by %s) [%s]\n", b.Title, b.Author, status)
	}
	fmt.Println("--------------------------")
}

func main() {
	// ตั้งชื่อห้องสมุด
	lib := Library{Name: "City Library"}

	// สร้างหนังสือ
	book1 := NewBook("Harry Potter", "J.K. Rowling")
	book2 := NewBook("The Stranger", "Albert Camus")

	// เอาเข้าห้องสมุด
	lib.AddBook(book1)
	lib.AddBook(book2)

	// ทดสอบยืม
	fmt.Println("\n>> User is borrowing Harry Potter...")
	lib.FindAndBorrow("Harry Potter")

	// ทดสอบยืมซ้ำ (ต้อง Error)
	fmt.Println("\n>> User is borrowing Harry Potter AGAIN...")
	lib.FindAndBorrow("Harry Potter")

	// ทดสอบยืมเล่มที่ไม่มี
	fmt.Println("\n>> User is borrowing Spiderman...")
	lib.FindAndBorrow("Spiderman")

	// ดูสรุปผล
	lib.ShowAllBooks()
}
