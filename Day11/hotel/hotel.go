package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// สร้าง Scanner ไว้รับค่า input จากคีย์บอร์ด (รองรับการเว้นวรรค)
var scanner = bufio.NewScanner(os.Stdin)

// Struct สำหรับเก็บข้อมูลห้องพัก
type room struct {
	roomID     int
	roomType   string
	price      float64
	guest      string
	isOccupied bool
}

// Slice สำหรับเก็บห้องทั้งหมด (เปรียบเสมือน Database)
var rooms []room

func main() {
	// --- 1. เริ่มต้นระบบ: สร้างห้องพัก 5 ห้อง ---
	for i := 1; i <= 5; i++ {
		// กำหนดค่าตัวแปรตามเงื่อนไข (ห้อง 1-3 เป็น Standard, 4-5 เป็น Suite)
		var rType string
		var rPrice float64

		if i <= 3 {
			rType = "Standard"
			rPrice = 1000
		} else {
			rType = "Suite"
			rPrice = 2500
		}

		// สร้างห้องและเพิ่มลงใน Slice (ใช้ ID 100+i เพื่อให้เป็น 101, 102...)
		newRoom := room{
			roomID:     100 + i,
			roomType:   rType,
			price:      rPrice,
			guest:      "",
			isOccupied: false,
		}
		rooms = append(rooms, newRoom)
	}

	fmt.Println("=== Mini Hotel System Initialized ===")

	// --- 2. Main Loop: วนลูปทำงานจนกว่าจะสั่งออก ---
	for {
		mainMenu()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
		scanner.Scan() // เคลียร์ Enter ที่ค้างใน Buffer

		switch choice {
		case 1:
			viewRooms()
		case 2:
			checkIn()
		case 3:
			checkOut()
		case 4:
			fmt.Println("Thank you! Goodbye.")
			return // จบโปรแกรม
		default:
			fmt.Println("Please enter a number between 1-4")
		}
	}
}

// ฟังก์ชันช่วยรับข้อความและตัดช่องว่าง
func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu() {
	fmt.Println("\n==== Main Menu ====")
	fmt.Println("1. View Rooms")
	fmt.Println("2. Check-in")
	fmt.Println("3. Check-out")
	fmt.Println("4. Exit")
}

func viewRooms() {
	fmt.Println("\n--- Room Status ---")
	for _, r := range rooms {
		if r.isOccupied {
			// ถ้ามีคนพัก: แสดงชื่อแขก
			fmt.Printf("[ Room %d ] %s (Occupied by %s)\n", r.roomID, r.roomType, r.guest)
		} else {
			// ถ้าว่าง: แสดงราคา
			fmt.Printf("[ Room %d ] %s (%.0f THB) - Empty\n", r.roomID, r.roomType, r.price)
		}
	}
	fmt.Println("-------------------")
}

func checkIn() {
	var inputID int
	fmt.Print("Enter Room ID: ")
	fmt.Scan(&inputID)
	scanner.Scan()

	found := false // ตัวแปรเช็คว่าเจอห้องไหม

	for i := range rooms {
		if rooms[i].roomID == inputID {
			found = true // ✅ เจอห้องแล้ว (ต้องยกธงตรงนี้ก่อนเช็คอย่างอื่น)

			if rooms[i].isOccupied {
				// ถ้าห้องไม่ว่าง
				fmt.Printf("Error: Room %d is already occupied by %s.\n", rooms[i].roomID, rooms[i].guest)
			} else {
				// ถ้าห้องว่าง -> รับชื่อแขกและเช็คอิน
				name := readLine("Enter Guest Name: ")
				updateIn(&rooms[i], name)
			}
			break // เจอแล้วหยุดหาทันที
		}
	}

	if !found {
		fmt.Println("Error: Room ID not found.")
	}
}

// ฟังก์ชันย่อยสำหรับบันทึกข้อมูล Check-in (รับ Pointer)
func updateIn(r *room, n string) {
	r.guest = n
	r.isOccupied = true
	fmt.Printf("✅ Check-in successful! Room %d assigned to %s.\n", r.roomID, r.guest)
}

func checkOut() {
	var inputID int
	fmt.Print("Enter Room ID: ")
	fmt.Scan(&inputID)
	scanner.Scan()

	found := false

	for i := range rooms {
		if rooms[i].roomID == inputID {
			found = true // ✅ เจอห้องแล้ว

			if rooms[i].isOccupied {
				// ถ้าห้องมีคนอยู่ -> ทำเรื่องคืนห้อง
				var nights float64
				fmt.Print("How many nights: ")
				fmt.Scan(&nights)
				scanner.Scan()

				updateOut(&rooms[i], nights)
			} else {
				// ถ้าห้องว่างอยู่แล้ว -> แจ้งเตือน
				fmt.Println("Error: This room is currently empty.")
			}
			break // เจอแล้วหยุดหา
		}
	}

	if !found {
		fmt.Println("Error: Room ID not found.")
	}
}

// ฟังก์ชันย่อยสำหรับบันทึกข้อมูล Check-out และคิดเงิน
func updateOut(r *room, n float64) {
	total := r.price * n
	fmt.Printf("💰 Total Bill: %.2f THB (Guest: %s)\n", total, r.guest)

	// ล้างข้อมูลห้องให้ว่าง
	r.guest = ""
	r.isOccupied = false
	fmt.Println("✅ Check-out complete. Room is now empty.")
}
