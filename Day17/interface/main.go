package main

import "fmt"

// Interface: ข้อตกลงร่วมกัน
type PaymentMethod interface {
	Pay(amount float64)
}

// Struct: เงินสด
type Cash struct {
	Balance float64
}

// Struct: บัตรเครดิต
type CreditCard struct {
	Balance float64
}

// Method สำหรับ Cash
func (c *Cash) Pay(amount float64) {
	if c.Balance >= amount {
		c.Balance -= amount
		// ปรับคำให้ดูเป็นทางการ: Payment Successful
		fmt.Printf("✅ Payment Successful: Paid %.2f THB (Cash) | Remaining Balance: %.2f THB\n", amount, c.Balance)
	} else {
		// ปรับคำแจ้งเตือน: Insufficient balance
		fmt.Printf("❌ Error: Insufficient cash balance (Need %.2f but have %.2f)\n", amount, c.Balance)
	}
}

// Method สำหรับ CreditCard
// เปลี่ยน receiver จาก 'r' เป็น 'c' (เพื่อให้สื่อถึง CreditCard)
func (c *CreditCard) Pay(amount float64) {
	if c.Balance >= amount {
		c.Balance -= amount
		// ปรับคำให้แยกออกว่าเป็นบัตรเครดิต
		fmt.Printf("💳 Payment Successful: Paid %.2f THB (Credit Card) | Remaining Limit: %.2f THB\n", amount, c.Balance)
	} else {
		fmt.Println("❌ Error: Insufficient credit limit")
	}
}

// ฟังก์ชันกลาง (Polymorphism)
func ProcessPayment(p PaymentMethod, amount float64) {
	fmt.Println("--------------------------------")
	fmt.Println("🔄 Processing payment...")
	p.Pay(amount) // เรียกใช้ Pay ของใครของมัน
}

func main() {
	myWallet := &Cash{Balance: 500}
	myCard := &CreditCard{Balance: 1000}

	// Case 1: จ่ายเงินสด (ผ่าน)
	ProcessPayment(myWallet, 350)

	// Case 2: จ่ายบัตรเครดิต (ผ่าน)
	ProcessPayment(myCard, 980)

	// Case 3: ลองจ่ายเกินตัว (ไม่ผ่าน)
	ProcessPayment(myWallet, 5000)
}
