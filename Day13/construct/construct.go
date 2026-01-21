package main

import "fmt"

type BankAccount struct {
	owner   string  // private
	balance float64 // private
}

// Constructor (Start with 'New...')
func NewAccount(o string, initialAmount float64) *BankAccount {
	if initialAmount < 0 {
		initialAmount = 0
	}
	return &BankAccount{
		owner:   o,
		balance: initialAmount,
	}
}

// 🛠️ แก้ชื่อเป็นตัวใหญ่ (Public Method) และแก้คำผิด Deposit
func (b *BankAccount) Deposit(amount float64) {
	if amount < 0 {
		fmt.Println("Error: Cannot deposit negative amount")
		return
	}
	b.balance += amount
	fmt.Printf("✅ Deposit: +%.2f | New Balance: %.2f THB\n", amount, b.balance)
}

// 🛠️ แก้ชื่อเป็นตัวใหญ่ (Public Method)
func (b *BankAccount) Withdraw(amount float64) {
	if amount > b.balance {
		fmt.Println("❌ Error: Insufficient funds")
		return
	}
	b.balance -= amount
	fmt.Printf("💸 Withdraw: -%.2f | New Balance: %.2f THB\n", amount, b.balance)
}

// 🛠️ Pro Tip: GetBalance ควร "Return ค่า" (Getter) ไม่ใช่แค่ปริ้น
// เพื่อให้คนอื่นเอาค่าตัวเลขไปคำนวณต่อได้
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

// เพิ่ม Method ใหม่สำหรับแสดงผลโดยเฉพาะ (Helper Method)
func (b *BankAccount) ShowStatement() {
	fmt.Println("------------------------------")
	fmt.Printf("Owner: %s\n", b.owner)
	fmt.Printf("Current Balance: %.2f THB\n", b.GetBalance()) // เรียกใช้ method ภายใน method ได้
	fmt.Println("------------------------------")
}

func main() {
	// สร้างบัญชี Phakin
	myAccount := NewAccount("Phakin", 500)
	myAccount.Deposit(450)
	myAccount.Deposit(2300)
	myAccount.Withdraw(1560)

	// เรียกดูยอดเงินแบบสวยงาม
	myAccount.ShowStatement()

	// 💡 ตัวอย่างประโยชน์ของ GetBalance ที่ return ค่า
	// สมมติเราอยากเช็คว่า "รวยหรือยัง?"
	if myAccount.GetBalance() > 1000 {
		fmt.Println(">>> You are rich! <<<")
	}

	fmt.Println("=================")

	// สร้างบัญชี Mom
	mom := NewAccount("Nipa", 500)
	mom.Deposit(40000)
	mom.Withdraw(3500)
	mom.ShowStatement()
}
