package character

import "fmt"

type Hero struct {
	Name  string
	hp    int
	power int
}

func NewHero(n string, h int, p int) *Hero { // รับค่า HP มาด้วยเลย จะได้กำหนดความอึดได้
	return &Hero{
		Name:  n,
		hp:    h,
		power: p,
	}
}

func (h *Hero) Attack(target *Hero) {
	target.hp -= h.power
	if target.hp < 0 {
		target.hp = 0
	}
	// เพิ่มบรรทัดนี้ เพื่อให้เห็นภาพการต่อสู้
	fmt.Printf("⚔️ %s attacks %s (Damage %d) -> %s HP remaining: %d\n",
		h.Name, target.Name, h.power, target.Name, target.hp)
}

func (h *Hero) GetHp() int {
	return h.hp
}
