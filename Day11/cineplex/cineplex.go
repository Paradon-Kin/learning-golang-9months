package main

import (
	"fmt"
	"strings" // เพิ่มเพื่อใช้ strings.ToUpper (แปลงตัวพิมพ์ใหญ่)
)

// สร้างโครงสร้างข้อมูลสำหรับ "ที่นั่ง"
type seat struct {
	row      string  // แถว (A, B)
	no       int     // เลขที่นั่ง (1-5)
	price    float64 // ราคา
	isBooked bool    // สถานะ: true=จองแล้ว, false=ว่าง
}

// ประกาศตัวแปร Global เพื่อให้ทุกฟังก์ชันมองเห็น
var seats []seat

func main() {
	fmt.Println("=== IMAX Theater Manager ===")

	// --- 1. ขั้นตอนการสร้างที่นั่ง (Generate Seats) ---
	rows := []string{"A", "B"}
	for _, r := range rows {
		for i := 1; i <= 5; i++ {
			// กำหนดราคา: แถว A แพงกว่า B
			var p float64
			if r == "A" {
				p = 300
			} else {
				p = 200
			}

			// สร้างที่นั่งใหม่
			newSeat := seat{
				row:      r,
				no:       i,
				price:    p,
				isBooked: false, // เริ่มต้นต้องว่างเสมอ
			}

			// เพิ่มเข้าใน Slice หลัก
			seats = append(seats, newSeat)
		}
	}

	// --- 2. Loop รอรับคำสั่ง (Menu Loop) ---
	for {
		mainMenu() // เรียกแสดงเมนู

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			view() // ดูผังที่นั่ง
		case 2:
			bookSeat() // จองที่นั่ง
		case 3:
			cancelSeat() // ยกเลิกที่นั่ง
		case 4:
			fmt.Println("Thank you. Goodbye!")
			return // จบการทำงาน
		default:
			fmt.Println("Please enter a number between 1-4") // แจ้งเตือนเมื่อกดผิด
		}
	}
}

func mainMenu() {
	fmt.Println("\n---- Main Menu ----")
	fmt.Println("1. View Seats")  // ดูที่นั่งทั้งหมด (เติม s)
	fmt.Println("2. Book Seat")   // จอง
	fmt.Println("3. Cancel Seat") // ยกเลิก
	fmt.Println("4. Exit")        // ออก
}

// ฟังก์ชันแสดงผลหน้าจอ
func view() {
	fmt.Println("\n--- Screen ---")
	for _, s := range seats {
		// เช็คสถานะ: ถ้าจองแล้วขึ้น XX, ถ้าว่างขึ้นเลขที่นั่ง
		if s.isBooked {
			fmt.Print("[ XX ] ")
		} else {
			fmt.Printf("[%s%d] ", s.row, s.no)
		}

		// Logic การขึ้นบรรทัดใหม่: ถ้าเป็นเก้าอี้เบอร์ 5 ให้ขึ้นบรรทัดใหม่
		if s.no == 5 {
			fmt.Println()
		}
	}
	fmt.Println("--------------")
}

// ฟังก์ชันรับค่าเพื่อทำการจอง
func bookSeat() {
	var bookRow string
	var bookNo int

	fmt.Print("Row (A/B): ")
	fmt.Scan(&bookRow)
	bookRow = strings.ToUpper(bookRow) // แปลงเป็นตัวใหญ่เสมอ (a -> A)

	fmt.Print("No (1-5): ")
	fmt.Scan(&bookNo)

	found := false // ตัวแปรเช็คว่าเจอที่นั่งไหม

	// วนลูปหาที่นั่งที่ตรงกับ Row และ No
	for i := range seats {
		if seats[i].row == bookRow && seats[i].no == bookNo {
			// เจอแล้ว! ส่ง Pointer (&) ไปให้ฟังก์ชัน book ทำงานต่อ
			book(&seats[i])
			found = true
			break // เจอแล้วหยุดหาทันที
		}
	}

	if !found {
		fmt.Println("Error: Seat not found.")
	}
}

// ฟังก์ชันเปลี่ยนสถานะเป็น "จอง" (รับ Pointer)
func book(s *seat) {
	if s.isBooked {
		// ถ้าจองไปแล้ว แจ้งเตือน
		fmt.Println("Error: Seat is already booked!")
	} else {
		// ถ้ายังไม่จอง เปลี่ยนเป็น true
		s.isBooked = true
		fmt.Println("Booking confirmed!")
	}
}

// ฟังก์ชันรับค่าเพื่อยกเลิก (Logic เหมือน bookSeat)
func cancelSeat() {
	var inputRow string
	var inputNo int

	fmt.Print("Row (A/B): ")
	fmt.Scan(&inputRow)
	inputRow = strings.ToUpper(inputRow)

	fmt.Print("No (1-5): ") // แก้ 1.5 เป็น 1-5
	fmt.Scan(&inputNo)

	found := false

	for i := range seats {
		if seats[i].row == inputRow && seats[i].no == inputNo {
			// เจอแล้ว! ส่ง Pointer (&) ไปให้ฟังก์ชัน cancel
			cancel(&seats[i])
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Error: Seat not found.")
	}
}

// ฟังก์ชันเปลี่ยนสถานะเป็น "ว่าง" (รับ Pointer)
func cancel(c *seat) {
	if !c.isBooked {
		// ถ้าที่นั่งมันว่างอยู่แล้ว จะยกเลิกไม่ได้
		fmt.Println("Error: Seat is already empty!")
	} else {
		// เปลี่ยนกลับเป็น false (ว่าง)
		c.isBooked = false
		fmt.Println("Seat cancelled successfully.")
	}
}
