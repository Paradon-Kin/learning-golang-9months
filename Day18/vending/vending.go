package main

import (
	"errors"
	"fmt"
)

var ErrSoldOut = errors.New("error: item sold out")

type Vending struct {
	Stock map[string]int
	Price map[string]int
}

func (v *Vending) Buy(item string, money int) error {
	stock, ok := v.Stock[item]
	if !ok {
		return fmt.Errorf("error: item '%s' not found", item)
	}

	// Step 2: Check stock
	if stock <= 0 {
		return ErrSoldOut
	}

	// Step 3: Check money
	itemPrice := v.Price[item]
	if money < itemPrice {
		diff := itemPrice - money
		return fmt.Errorf("error: insufficient funds, need %d THB more", diff)
	}

	// Step 4: Deduct stock
	v.Stock[item]-- // เขียนย่อจาก -= 1
	return nil
}

func main() {
	machine := &Vending{
		Stock: map[string]int{
			"Coke":  20,
			"Pepsi": 13,
			"Water": 0, // เพิ่มตัวอย่างของหมด
		},
		Price: map[string]int{
			"Coke":  15,
			"Pepsi": 16,
			"Water": 10,
		},
	}

	// Case 1: ซื้อสำเร็จ
	fmt.Println("\n1. Buying Coke (Pay 15)...")
	err1 := machine.Buy("Coke", 15)
	handleTransaction(err1)

	// Case 2: เงินไม่พอ
	fmt.Println("\n2. Buying Pepsi (Pay 10)...")
	err2 := machine.Buy("Pepsi", 10)
	handleTransaction(err2)

	// Case 3: ของหมด
	fmt.Println("\n3. Buying Water (Pay 20)...")
	err3 := machine.Buy("Water", 20)
	handleTransaction(err3)
}

// สร้างฟังก์ชันแยก เพื่อลดโค้ดซ้ำใน main (Optional)
func handleTransaction(err error) {
	if err == nil {
		fmt.Println("✅ Success: Enjoy your drink!")
		return
	}

	// เช็ค Error เฉพาะเจาะจง
	if err == ErrSoldOut {
		fmt.Println("❌ Sorry, this item is sold out.")
	} else {
		// Error อื่นๆ (เช่น เงินไม่พอ, หาของไม่เจอ)
		fmt.Println("⚠️", err)
	}
}
