package main

import (
	"fmt"
)

// Struct สำหรับที่นั่ง
type Seat struct {
	no       int
	price    float64
	isBooked bool   // true = จองแล้ว, false = ว่าง
	owner    string // ชื่อคนจอง
}

// Struct สำหรับงานคอนเสิร์ต
type Concert struct {
	Title string
	Seats []*Seat // เก็บ Pointer ของที่นั่ง
}

// Constructor: สร้างงานคอนเสิร์ตและเสกที่นั่งมารอไว้เลย
func NewConcert(title string, totalSeats int, price float64) *Concert {
	// 1. สร้าง Object Concert ว่างๆ
	c := &Concert{
		Title: title,
		Seats: []*Seat{},
	}

	// 2. วนลูปสร้างที่นั่งตามจำนวนที่ระบุ (totalSeats)
	for i := 1; i <= totalSeats; i++ {
		newSeat := &Seat{
			no:       i,
			price:    price,
			isBooked: false, // เริ่มต้นต้อง "ว่าง"
			owner:    "",    // ยังไม่มีเจ้าของ
		}
		// 3. ยัดที่นั่งใส่เข้าไปใน Concert
		c.Seats = append(c.Seats, newSeat)
	}

	return c
}

// Method จองตั๋ว
func (c *Concert) BookTicket(seatNo int, name string) {
	found := false

	for i := range c.Seats {
		// ค้นหาเลขที่นั่งที่ตรงกัน
		if seatNo == c.Seats[i].no {
			found = true

			// เช็คว่าที่นั่งว่างไหม? (ถ้า isBooked เป็น false คือว่าง)
			if !c.Seats[i].isBooked {
				// --- เริ่มการจอง ---
				c.Seats[i].owner = name
				c.Seats[i].isBooked = true // 🛠️ สำคัญ: จองแล้วต้องเปลี่ยนเป็น TRUE

				fmt.Printf("✅ Success: Seat %d reserved for '%s'.\n", seatNo, name)
			} else {
				// ถ้าไม่ว่าง
				fmt.Printf("❌ Error: Seat %d is already occupied by '%s'.\n", seatNo, c.Seats[i].owner)
			}
			break // เจอแล้วหยุดหาทันที
		}
	}

	if !found {
		fmt.Printf("⚠️ Error: Seat number %d not found.\n", seatNo)
	}
}

// Method ยกเลิกตั๋ว
func (c *Concert) CancelTicket(seatNo int) {
	found := false

	for i := range c.Seats {
		if seatNo == c.Seats[i].no {
			found = true

			// เช็คก่อนว่ามันถูกจองอยู่จริงไหม (กันพลาด)
			if c.Seats[i].isBooked {
				// --- เริ่มการยกเลิก ---
				oldOwner := c.Seats[i].owner // เก็บชื่อไว้โชว์ก่อนลบ
				c.Seats[i].isBooked = false  // 🛠️ สำคัญ: ยกเลิกแล้วต้องกลับมาเป็น FALSE (ว่าง)
				c.Seats[i].owner = ""        // ลบชื่อเจ้าของออก

				fmt.Printf("🗑️ Cancelled: Seat %d (was owned by %s) is now empty.\n", seatNo, oldOwner)
			} else {
				fmt.Printf("⚠️ Warning: Seat %d is already empty.\n", seatNo)
			}
			break
		}
	}

	if !found {
		fmt.Printf("⚠️ Error: Seat number %d not found.\n", seatNo)
	}
}

// Method สรุปยอดขาย
func (c *Concert) ShowSummary() {
	var totalSales float64 = 0.0
	var availableSeats int = 0 // เปลี่ยนชื่อจาก s เป็น availableSeats ให้อ่านง่าย

	for _, seat := range c.Seats {
		if seat.isBooked {
			// ถ้าจองแล้ว -> บวกเงินเข้ายอดขาย
			totalSales += seat.price
		} else {
			// ถ้ายังไม่จอง -> นับเป็นที่ว่าง
			availableSeats++
		}
	}

	totalSeats := len(c.Seats)

	fmt.Println("\n==============================")
	fmt.Printf("🎤 Concert: %s\n", c.Title)
	fmt.Println("==============================")
	fmt.Printf("💰 Total Revenue : %.2f THB\n", totalSales)
	fmt.Printf("🪑 Availability  : %d / %d seats\n", availableSeats, totalSeats)
	fmt.Println("==============================")
}

func main() {
	// สร้างคอนเสิร์ต EXO มี 10 ที่นั่ง ราคาใบละ 2500
	concert := NewConcert("EXO Planet #5", 10, 2500)

	// ทดสอบการจอง
	concert.BookTicket(4, "Phakin")
	concert.BookTicket(2, "Jack")
	concert.BookTicket(8, "Game")
	concert.BookTicket(5, "Jame")

	// ทดสอบจองซ้ำ (ต้องขึ้น Error)
	concert.BookTicket(4, "Som")

	// ทดสอบยกเลิก
	fmt.Println("--- Canceling Ticket ---")
	concert.CancelTicket(8) // ยกเลิกของ Game

	// แสดงสรุปผล
	concert.ShowSummary()
}
