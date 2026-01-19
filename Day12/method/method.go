package main

import (
	"fmt"
)

type hero struct {
	name  string
	hp    int
	power int
}

// ไม่ต้องเช็ค h.hp <= 0 ตรงนี้ เพราะคนโจมตีเลือดไม่ลด
func (h *hero) attack(target *hero) {
	target.hp -= h.power

	// กันเลือดติดลบ
	if target.hp < 0 {
		target.hp = 0
	}

	fmt.Printf("⚔️ %s attacks %s! (Damage %d) -> %s HP: %d\n",
		h.name, target.name, h.power, target.name, target.hp)
}

func main() {
	// ปรับเลือดให้สูสีกันหน่อย
	myhero := hero{name: "Iron Man", hp: 400, power: 120}
	target := hero{name: "Thanos", hp: 500, power: 80}

	fmt.Println("=== 🔔 FIGHT START! 🔔 ===")

	// Game Loop: สู้กันจนกว่าจะมีคนตาย
	for myhero.hp > 0 && target.hp > 0 {
		// 1. เราตีเขาก่อน
		myhero.attack(&target)
		if target.hp == 0 {
			break // ถ้าเขาตาย หยุดทันที
		}

		// 2. เขาตีสวน
		target.attack(&myhero)
		if myhero.hp == 0 {
			break // ถ้าเราตาย หยุดทันที
		}

		fmt.Println("...") // เว้นวรรคนิดนึง
	}

	fmt.Println("==========================")

	// เช็คหาผู้ชนะ (ใครที่เลือดยังเหลืออยู่ คือผู้ชนะ)
	if myhero.hp > 0 {
		fmt.Printf("🏆 The Winner is %s!\n", myhero.name)
	} else {
		fmt.Printf("💀 The Winner is %s!\n", target.name)
	}
}
