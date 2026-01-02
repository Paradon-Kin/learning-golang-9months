package main

import "fmt"

// Exchange rate constants (as of January 2025)
// ค่าคงที่อัตราแลกเปลี่ยน (ณ มกราคม 2025)
const (
	THBtoUSD = 0.029 // 1 THB = 0.029 USD
	THBtoJPY = 4.32  // 1 THB = 4.32 JPY
	THBtoEUR = 0.027 // 1 THB = 0.027 EUR

	USDtoTHB = 34.48 // 1 USD = 34.48 THB
	JPYtoTHB = 0.23  // 1 JPY = 0.23 THB
	EURtoTHB = 37.04 // 1 EUR = 37.04 THB
)

// Exchange fee as percentage (1%)
// ค่าธรรมเนียมเป็นเปอร์เซ็นต์ (1%)
const FeePercentage = 0.01

func main() {
	// Display welcome message
	// แสดงข้อความต้อนรับ
	fmt.Println("======================================")
	fmt.Println("  เครื่องแปลงสกุลเงิน (มีค่าธรรมเนียม) ")
	fmt.Printf("         ค่าธรรมเนียม: %.0f%%\n", FeePercentage*100)
	fmt.Println("======================================")
	fmt.Println()

	// Step 1: Choose conversion direction
	// ขั้นตอนที่ 1: เลือกทิศทางการแปลง
	fmt.Println("📍 Choose conversion direction:")
	fmt.Println("1. THB → Foreign Currency (บาท → เงินต่างประเทศ)")
	fmt.Println("2. Foreign Currency → THB (เงินต่างประเทศ → บาท)")

	var direction int
	fmt.Print("\nYour choice (1 or 2): ")
	fmt.Scan(&direction)

	// Validate direction choice
	// ตรวจสอบตัวเลือกทิศทาง
	if direction != 1 && direction != 2 {
		fmt.Println("❌ Invalid choice. Please enter 1 or 2")
		return
	}

	fmt.Println()

	// Branch based on direction
	// แยกทางตามทิศทาง
	if direction == 1 {
		convertFromTHB()
	} else {
		convertToTHB()
	}
}

// convertFromTHB converts Thai Baht to foreign currency
// convertFromTHB แปลงบาทไทยเป็นเงินต่างประเทศ
func convertFromTHB() {
	var thbAmount float64

	// Step 2: Choose target currency
	// ขั้นตอนที่ 2: เลือกสกุลเงินปลายทาง
	fmt.Println("💱 เลือกสกุลเงินที่จะแปลง:")
	fmt.Println("1. USD - US Dollar (ดอลลาร์สหรัฐ)")
	fmt.Println("2. JPY - Japanese Yen (เยนญี่ปุ่น)")
	fmt.Println("3. EUR - Euro (ยูโร)")

	var currencyChoice int
	fmt.Print("\nเลือก (1-3): ")
	fmt.Scan(&currencyChoice)

	// Validate currency choice
	// ตรวจสอบตัวเลือกสกุลเงิน
	if currencyChoice < 1 || currencyChoice > 3 {
		fmt.Println("❌ กรุณาใส่ตัวเลข 1-3")
		return
	}

	// Set currency details based on choice
	// กำหนดรายละเอียดสกุลเงินตามที่เลือก
	var currencyName string
	var currencySymbol string
	var exchangeRate float64

	switch currencyChoice {
	case 1:
		currencyName = "USD"
		currencySymbol = "$"
		exchangeRate = THBtoUSD
	case 2:
		currencyName = "JPY"
		currencySymbol = "¥"
		exchangeRate = THBtoJPY
	case 3:
		currencyName = "EUR"
		currencySymbol = "€"
		exchangeRate = THBtoEUR
	}

	// Step 3: Get THB amount from user
	// ขั้นตอนที่ 3: รับจำนวนเงินบาทจากผู้ใช้
	fmt.Print("\nกรอกเงิน (THB): ")
	fmt.Scan(&thbAmount)

	// Validate amount
	// ตรวจสอบจำนวนเงิน
	if thbAmount <= 0 {
		fmt.Println("❌ จำนวนเงินต้องมากกว่า 0")
		return
	}

	// Calculate fee and final amount
	// คำนวณค่าธรรมเนียมและยอดสุดท้าย
	fee := thbAmount * FeePercentage
	thbAfterFee := thbAmount - fee
	convertedAmount := thbAfterFee * exchangeRate

	// Display results
	// แสดงผลลัพธ์
	fmt.Println("\n========== ผลการแปลง ==========")
	fmt.Printf("เงินต้น: %.2f THB\n", thbAmount)
	fmt.Printf("ค่าธรรมเนียม (%.0f%%): -%.2f THB\n", FeePercentage*100, fee)
	fmt.Printf("เงินหลังหักค่าธรรมเนียม: %.2f THB\n", thbAfterFee)
	fmt.Printf("💰 แปลงได้: %s%.2f %s\n", currencySymbol, convertedAmount, currencyName)
	fmt.Printf("อัตราแลกเปลี่ยน: 1 THB = %.4f %s\n", exchangeRate, currencyName)
	fmt.Println("===================================")
}

// convertToTHB converts foreign currency to Thai Baht
// convertToTHB แปลงเงินต่างประเทศเป็นบาทไทย
func convertToTHB() {
	var foreignAmount float64

	// Step 2: Choose source currency
	// ขั้นตอนที่ 2: เลือกสกุลเงินต้นทาง
	fmt.Println("💱 Choose source currency:")
	fmt.Println("1. USD - US Dollar (ดอลลาร์สหรัฐ)")
	fmt.Println("2. JPY - Japanese Yen (เยนญี่ปุ่น)")
	fmt.Println("3. EUR - Euro (ยูโร)")

	var currencyChoice int
	fmt.Print("\nYour choice (1-3): ")
	fmt.Scan(&currencyChoice)

	// Validate currency choice
	// ตรวจสอบตัวเลือกสกุลเงิน
	if currencyChoice < 1 || currencyChoice > 3 {
		fmt.Println("❌ Invalid choice. Please enter 1-3")
		return
	}

	// Set currency details based on choice
	// กำหนดรายละเอียดสกุลเงินตามที่เลือก
	var currencyName string
	var currencySymbol string
	var exchangeRate float64 // ✅ FIXED: consistent naming

	switch currencyChoice {
	case 1:
		currencyName = "USD"
		currencySymbol = "$"
		exchangeRate = USDtoTHB
	case 2:
		currencyName = "JPY"
		currencySymbol = "¥"
		exchangeRate = JPYtoTHB
	case 3:
		currencyName = "EUR"
		currencySymbol = "€"
		exchangeRate = EURtoTHB
	}

	// Step 3: Get foreign currency amount from user
	// ขั้นตอนที่ 3: รับจำนวนเงินต่างประเทศจากผู้ใช้
	fmt.Printf("\nกรอกจำนวนเงิน %s: ", currencyName)
	fmt.Scan(&foreignAmount)

	// Validate amount
	// ตรวจสอบจำนวนเงิน
	if foreignAmount <= 0 {
		fmt.Println("❌ จำนวนเงินต้องมากกว่า 0")
		return
	}

	// Calculate conversion and fee
	// คำนวณการแปลงและค่าธรรมเนียม
	thbBeforeFee := foreignAmount * exchangeRate
	fee := thbBeforeFee * FeePercentage
	thbAfterFee := thbBeforeFee - fee

	// Display results
	// แสดงผลลัพธ์
	fmt.Println("\n========== ผลการแปลง ==========")
	fmt.Printf("เงินต้น: %s%.2f %s\n", currencySymbol, foreignAmount, currencyName)
	fmt.Printf("แปลงได้: %.2f THB\n", thbBeforeFee)
	fmt.Printf("ค่าธรรมเนียม (%.0f%%): -%.2f THB\n", FeePercentage*100, fee)
	fmt.Printf("💰 ยอดสุดท้าย: %.2f THB\n", thbAfterFee)
	fmt.Printf("อัตราแลกเปลี่ยน: 1 %s = %.2f THB\n", currencyName, exchangeRate)
	fmt.Println("===================================")
}
