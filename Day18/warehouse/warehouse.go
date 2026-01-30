package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("item not found")

type Warehouse struct {
	Stock map[string]int
}

func (w *Warehouse) PickItem(name string, amount int) error {
	currentAmount, ok := w.Stock[name]
	if !ok {
		return ErrorNotFound
	} else if currentAmount < amount {
		return fmt.Errorf("insufficient stock for %s: need %d but have %d", name, amount, currentAmount)
	}

	w.Stock[name] -= amount
	return nil
}

func main() {
	w := &Warehouse{
		Stock: map[string]int{
			"Cola":  10,
			"Water": 5,
		},
	}

	// --- Case 1: ของไม่พอ (Dynamic Error) ---
	fmt.Println("1. Buying 20 Cola...")
	err1 := w.PickItem("Cola", 20)
	if err1 != nil {
		fmt.Println(err1) // จะปริ้นว่ามี 10 แต่ต้องการ 20
	}

	// --- Case 2: ไม่มีของ (Static Error) ---
	fmt.Println("\n2. Buying Beer...")
	err2 := w.PickItem("Beer", 1)
	if err2 == ErrorNotFound {
		fmt.Println("Error: We don't sell this item.")
	} else if err2 != nil {
		fmt.Println(err2)
	}

	// --- Case 3: สำเร็จ ---
	fmt.Println("\n3. Buying 2 Water...")
	err3 := w.PickItem("Water", 2)
	if err3 == nil {
		fmt.Printf("✅ Success! Water left: %d\n", w.Stock["Water"])
	}
}
