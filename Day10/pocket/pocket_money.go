package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// สร้าง Scanner ไว้รับค่า input จากคีย์บอร์ด (ใช้ bufio เพื่อแก้ปัญหาการรับ string ยาวๆ)
var scanner = bufio.NewScanner(os.Stdin)

// Struct สำหรับเก็บข้อมูลธุรกรรมการเงิน 1 รายการ
type transaction struct {
	id     int
	class  string // เก็บประเภท: "income" (รายรับ) หรือ "expense" (รายจ่าย)
	amount float64
	memo   string
}

// Slice สำหรับเก็บรายการ transaction ทั้งหมด (Database จำลองในแรม)
var transactions []transaction

func main() {
	fmt.Println("=========== Pocket Money ===========")

	// โหลดข้อมูลตัวอย่าง
	sampleData()

	// Loop หลักของโปรแกรม (เพื่อให้เมนูแสดงวนไปเรื่อยๆ จนกว่าจะสั่งออก)
	for {
		mainMenu()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		scanner.Scan() // เคลียร์ปุ่ม Enter ที่ค้างอยู่ใน Buffer (สำคัญมาก)
		fmt.Println()

		switch choice {
		case 1:
			addTransaction() // เปลี่ยนชื่อฟังก์ชันให้ตรงกับการทำงานจริง
		case 2:
			editTransaction()
		case 3:
			viewSummary()
		case 4:
			viewAll()
		case 5:
			fmt.Println("Thank you. Goodbye!")
			return // จบการทำงานของฟังก์ชัน main (ปิดโปรแกรม)
		default:
			fmt.Println("Please enter a number between 1-5")
		}
	}
}

// ฟังก์ชันสำหรับรับข้อความ (String) และตัดช่องว่างหัวท้ายออก
func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

// สร้างข้อมูลจำลองใส่ลงไปใน Slice
func sampleData() {
	transactions = []transaction{
		{id: 1, class: "income", amount: 360, memo: "Dad"},
		{id: 2, class: "expense", amount: 45, memo: "Chicken Rice"},
		{id: 3, class: "income", amount: 340, memo: "Mom"},
	}
}

func mainMenu() {
	fmt.Println("\n---- Main Menu -----")
	fmt.Println("1. Add Income/Expense")
	fmt.Println("2. Edit Transaction")
	fmt.Println("3. View Balance")
	fmt.Println("4. View All Transactions")
	fmt.Println("5. Exit")
}

// ฟังก์ชันเพิ่มรายการ (รับทั้งรายรับและรายจ่าย)
func addTransaction() {
	var t transaction

	// Logic สร้าง ID อัตโนมัติ: ถ้าไม่มีข้อมูลเริ่มที่ 1, ถ้ามีให้เอาตัวสุดท้าย +1
	if len(transactions) == 0 {
		t.id = 1
	} else {
		t.id = transactions[len(transactions)-1].id + 1
	}

	fmt.Printf("Transaction ID: %d\n", t.id)

	t.class = readLine("Type (income/expense): ")

	fmt.Print("Amount: ")
	fmt.Scan(&t.amount)
	scanner.Scan() // เคลียร์ Enter หลังรับตัวเลข

	// ตรวจสอบความถูกต้องของตัวเลข (ต้องไม่ติดลบหรือเป็น 0)
	if t.amount < 1 {
		fmt.Println("Invalid amount. Please try again.")
		return
	}

	t.memo = readLine("Memo: ")

	// บันทึกข้อมูลลงใน Slice
	transactions = append(transactions, t)
	fmt.Printf("Added %s successfully.\n", t.class)
}

// ฟังก์ชันแก้ไขรายการ (ใช้ Pointer เพื่อแก้ค่าใน Memory โดยตรง)
func editTransaction() {
	var searchID int
	fmt.Print("Enter ID to edit: ")
	fmt.Scan(&searchID)
	scanner.Scan() // เคลียร์ Enter

	found := false

	// วนลูปหา ID ที่ต้องการแก้ไข
	for i, t := range transactions {
		if t.id == searchID {
			found = true

			// แสดงข้อมูลเดิมให้เห็นก่อนแก้
			fmt.Printf("ID: %d\n", t.id)
			fmt.Printf("Type: %s\n", t.class)
			fmt.Printf("Current Amount: %.2f\n", t.amount)
			fmt.Printf("Memo: %s\n", t.memo)

			var newAmount float64
			fmt.Print("Enter new amount: ")
			fmt.Scan(&newAmount)
			scanner.Scan() // เคลียร์ Enter

			// ส่ง Address (&) ของข้อมูลตัวที่ i ไปให้ฟังก์ชันทำการแก้ไข
			editAmount(&transactions[i], newAmount)

			// เจอแล้วให้หยุดหาทันทีเพื่อประหยัดเวลา (Optimization)
			break
		}
	}

	if !found {
		fmt.Println("Transaction ID not found.")
	}
}

// ฟังก์ชันย่อย รับ Pointer มาเพื่อแก้ค่า amount
func editAmount(a *transaction, n float64) {
	a.amount = n
	fmt.Println("Amount updated successfully.")
}

// ฟังก์ชันคำนวณและสรุปยอดเงิน
func viewSummary() {
	var totalIncome float64
	var totalExpense float64

	for _, s := range transactions {
		// ใช้ EqualFold เพื่อเปรียบเทียบแบบไม่สนใจตัวพิมพ์เล็ก/ใหญ่ (income == Income)
		if strings.EqualFold(s.class, "income") {
			totalIncome += s.amount
		} else {
			totalExpense += s.amount
		}
	}

	balance := totalIncome - totalExpense
	fmt.Println("---- Summary Report ----")
	fmt.Printf("Total Income:  %10.2f\n", totalIncome)
	fmt.Printf("Total Expense: %10.2f\n", totalExpense)
	fmt.Println("------------------------")
	fmt.Printf("Balance:       %10.2f\n", balance)
	fmt.Println("------------------------")
}

// ฟังก์ชันแสดงประวัติทั้งหมด
func viewAll() {
	fmt.Println("---- All Transactions ----")
	for _, v := range transactions {
		fmt.Printf("ID: %d\n", v.id)
		fmt.Printf("Type: %s\n", v.class)
		fmt.Printf("Amount: %.2f\n", v.amount)
		fmt.Printf("Memo: %s\n", v.memo)
		fmt.Println("--------------------------")
	}
}
