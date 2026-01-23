package main

import (
	"fmt"
	"my-game/character"
	"time" // แถม: เอาไว้หน่วงเวลาให้ดูตื่นเต้น
)

func main() {
	fmt.Println("==== 🥊 BATTLE START 🥊 ====")

	// สร้าง Hero (ใส่เลือดเยอะหน่อย จะได้สู้นานๆ)
	myHero := character.NewHero("Dr.Strange", 100, 20)
	enemy := character.NewHero("Thanos", 150, 15) // Thanos เลือดเยอะกว่า แต่ตีน้อยกว่า

	// --- Game Loop ---
	// วนลูปตราบใดที่ "ทั้งคู่" ยังมีชีวิตอยู่
	for myHero.GetHp() > 0 && enemy.GetHp() > 0 {

		// 1. ฝั่งเราโจมตี
		myHero.Attack(enemy)
		if enemy.GetHp() == 0 {
			fmt.Println("--------------------------")
			fmt.Printf("🏆 WINNER IS: %s !!\n", myHero.Name)
			break // จบลูปทันที
		}

		time.Sleep(500 * time.Millisecond) // (แถม) หยุดรอ 0.5 วินาที ให้ดูสมจริง

		// 2. ฝั่งศัตรูสวนคืน
		enemy.Attack(myHero)
		if myHero.GetHp() == 0 {
			fmt.Println("--------------------------")
			fmt.Printf("💀 GAME OVER! Winner is: %s\n", enemy.Name)
			break // จบลูปทันที
		}

		time.Sleep(500 * time.Millisecond) // (แถม) หยุดรอ 0.5 วินาที
		fmt.Println("...")
	}

	fmt.Println("==== BATTLE END ====")
}
