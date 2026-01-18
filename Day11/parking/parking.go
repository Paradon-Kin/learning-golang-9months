package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// สร้าง Scanner ไว้รับค่า Input (รองรับการเว้นวรรค)
var scanner = bufio.NewScanner(os.Stdin)

// Struct สำหรับช่องจอดรถแต่ละช่อง
type slot struct {
	id         int
	plate      string
	hours      int
	isOccupied bool // true = มีรถจอด, false = ว่าง
}

// Slice เก็บข้อมูลช่องจอดทั้งหมด (Database ในแรม)
var slots []slot

func main() {
	// --- 1. เริ่มต้นระบบ: สร้างช่องจอดรถว่างๆ 5 ช่อง ---
	for i := 1; i <= 5; i++ {
		newSlot := slot{
			id:         i,
			plate:      "",
			hours:      0,
			isOccupied: false,
		}
		slots = append(slots, newSlot)
	}

	fmt.Println("=== Smart Parking System Initialized ===")

	// --- 2. Main Loop: วนลูปทำงานจนกว่าจะสั่งออก ---
	for {
		mainMenu()

		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
		scanner.Scan() // 🔥 สำคัญมาก: เคลียร์ปุ่ม Enter ที่ค้างอยู่ใน Buffer

		switch choice {
		case 1:
			viewParking()
		case 2:
			parkCar()
		case 3:
			leaveCar()
		case 4:
			fmt.Println("Thank you! Goodbye.")
			return // จบโปรแกรม
		default:
			fmt.Println("Invalid input. Please enter a number between 1-4.")
		}
	}
}

// ฟังก์ชันช่วยรับค่า String และตัดช่องว่าง
func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func mainMenu() {
	fmt.Println("\n==== Main Menu ====")
	fmt.Println("1. View Parking Status")
	fmt.Println("2. Park Car (Check-in)")
	fmt.Println("3. Leave Car (Check-out)")
	fmt.Println("4. Exit")
}

func viewParking() {
	fmt.Println("\n--- Parking Status ---")
	for _, v := range slots {
		if v.isOccupied {
			// ถ้ามีรถจอด ให้โชว์ทะเบียนและชั่วโมง
			fmt.Printf("[ Slot %d: %s (%d hrs) ]\n", v.id, v.plate, v.hours)
		} else {
			// ถ้าว่าง ให้โชว์ Empty
			fmt.Printf("[ Slot %d: Empty ]\n", v.id)
		}
	}
	fmt.Println("----------------------")
}

func parkCar() {
	// รับข้อมูลรถเข้า
	plate := readLine("Enter Plate Number: ")

	var hours int
	fmt.Print("Enter Parking Hours: ")
	fmt.Scan(&hours)
	scanner.Scan() // เคลียร์ Enter

	found := false // ตัวแปรเช็คว่าเจอช่องว่างไหม

	// วนลูปหา "ช่องว่างช่องแรก"
	for i := range slots {
		if !slots[i].isOccupied { // เช็คว่าช่องนี้ว่างไหม (!false คือ true)
			// ถ้าว่าง -> บันทึกข้อมูลลงไป
			slots[i].plate = plate
			slots[i].hours = hours
			slots[i].isOccupied = true

			fmt.Printf("✅ Success! Parked at Slot %d\n", slots[i].id)
			found = true
			break // 🛑 หยุดหาทันที (ไม่งั้นจะจอดรถคันเดียวทุกช่อง!)
		}
	}

	// ถ้าวนครบทุกช่องแล้วยังไม่เจอที่ว่าง
	if !found {
		fmt.Println("❌ Parking Full! No slots available.")
	}
}

func leaveCar() {
	inputPlate := readLine("Enter Plate to Leave: ")

	found := false // ตัวแปรเช็คว่าเจอรถไหม

	// วนลูปหารถตามทะเบียน
	for i := range slots {
		// ใช้ EqualFold เพื่อเทียบแบบไม่สนตัวพิมพ์เล็กใหญ่ (กข123 == กข123)
		if strings.EqualFold(slots[i].plate, inputPlate) {

			// เจอแล้ว! ส่ง Pointer ไปให้ฟังก์ชัน clear จัดการต่อ
			clear(&slots[i])

			found = true
			break // 🛑 เจอแล้วหยุดหา
		}
	}

	// ถ้าหาไม่เจอ
	if !found {
		fmt.Println("❌ Error: Car not found!")
	}
}

// ฟังก์ชันเคลียร์ช่องจอด (รับ Pointer เพื่อแก้ไขค่าจริง)
func clear(s *slot) {
	// 1. 💰 คิดเงิน "ก่อน" ลบข้อมูล (สำคัญมาก!)
	// สมมติชั่วโมงละ 20 บาท
	total := s.hours * 20
	fmt.Printf("💰 Total Bill: %d THB\n", total)

	// 2. 🧹 ล้างข้อมูลช่องจอด
	s.isOccupied = false
	s.plate = "" // ต้องลบทะเบียนทิ้งด้วย ไม่งั้นข้อมูลเก่าค้าง
	s.hours = 0

	fmt.Println("✅ Check-out complete. Slot is now empty.")
}
