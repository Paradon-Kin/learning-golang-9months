package main

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

type vendingMachine struct {
	coins     int
	inventory []product // ❌ ลบ () ออกครับ เขียนแค่นี้พอ
}

func (vm *vendingMachine) addProduct(name string, price int, stock int) {
	newProduct := product{
		name:  name,
		price: price,
		stock: stock,
	}
	vm.inventory = append(vm.inventory, newProduct)
	fmt.Printf("Added: %s (Price: %d, Stock: %d)\n", name, price, stock)
}

func (vm *vendingMachine) insertCoin(amount int) {
	vm.coins += amount
	fmt.Printf("Inserted coin: %d | Current Balance: %d\n", amount, vm.coins)
}

func (vm *vendingMachine) selectProduct(name string) {
	// 🔥 จุดสำคัญ: วนลูปหาของ
	for i := range vm.inventory {
		// เช็คชื่อสินค้า
		if name == vm.inventory[i].name {

			// เจอแล้ว! เช็คสต็อกต่อ
			if vm.inventory[i].stock > 0 {

				// เช็คเงินต่อ
				if vm.coins >= vm.inventory[i].price {
					// เงินพอ -> ตัดเงิน, ตัดของ
					vm.coins -= vm.inventory[i].price
					vm.inventory[i].stock--
					fmt.Printf("✅ Dispensing [%s]. Enjoy!\n", name)
				} else {
					fmt.Println("❌ Not enough money")
				}
			} else {
				fmt.Println("❌ Out of stock")
			}

			return // 🛑 เจอสินค้าแล้ว (ไม่ว่าจะซื้อได้หรือไม่ได้) ให้จบฟังก์ชันทันที!
		}
	}

	// ถ้าวนลูปจนจบแล้วยังไม่เจอ 'return' ข้างบน แสดงว่าไม่มีสินค้านี้
	fmt.Println("❌ Product not found")
}

func (vm *vendingMachine) returnChange() {
	change := vm.coins
	vm.coins = 0 // เคลียร์เงินในตู้
	fmt.Printf("💰 Returning change: %d\n", change)
}

func main() {
	vm := vendingMachine{}

	// เติมของ
	vm.addProduct("Coke", 10, 30)
	vm.addProduct("Pepsi", 20, 15)
	vm.addProduct("Ishitan", 21, 23)
	fmt.Println("----------------")

	// ทดสอบซื้อของ
	vm.insertCoin(15)        // หยอด 15 บาท
	vm.selectProduct("Coke") // ซื้อโค้ก (10 บาท) -> ซื้อได้ เหลือ 5 บาท

	vm.returnChange() // คืนเงิน (ควรได้คืน 5 บาท)
}
