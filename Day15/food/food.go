package main

import (
	"fmt"
)

type Order struct {
	Id       int
	FoodName string
	Price    float64
	status   string // private: ใช้ตัวเล็กตามโจทย์
}

type Rider struct {
	Name    string
	wallet  float64 // private: เงินในกระเป๋า ก็ไม่ควรให้ใครมาแก้ตรงๆ
	History []*Order
}

// Constructor Order
func NewOrder(id int, name string, price float64) *Order {
	return &Order{
		Id:       id,
		FoodName: name,
		Price:    price,
		status:   "Pending", // เริ่มต้นรออาหาร
	}
}

// Constructor Rider
func NewRider(name string) *Rider {
	return &Rider{
		Name:    name,
		wallet:  0.00,
		History: []*Order{},
	}
}

// Helper Method: เอาไว้ดึงสถานะไปโชว์ (เพราะ status เป็น private)
func (o *Order) GetStatus() string {
	return o.status
}

// ขั้นตอนที่ 1: รับงาน (ยังไม่ได้เงิน)
func (r *Rider) AcceptJob(o *Order) {
	// เปลี่ยนสถานะออเดอร์
	o.status = "On Delivery"

	// เก็บ Pointer ใบงานนี้ลงประวัติ
	r.History = append(r.History, o)

	fmt.Printf("🛵 Rider %s accepted job: %s (Status: %s)\n", r.Name, o.FoodName, o.status)
}

// ขั้นตอนที่ 2: ส่งงานเสร็จ (รับเงินตอนนี้)
func (r *Rider) FinishJob(o *Order) {
	// เปลี่ยนสถานะเป็นส่งเสร็จแล้ว
	o.status = "Delivered"

	// คำนวณเงิน (หัก GP 20%)
	netIncome := o.Price * 0.80 // คูณ 0.80 คือเหลือ 80% (คิดแบบย่อ)

	// เอาเงินเข้ากระเป๋า (ใช้ += สะสมเงิน)
	r.wallet += netIncome

	fmt.Printf("✅ Delivered! Earned: %.2f THB (Wallet Updated)\n", netIncome)
	fmt.Println("------------------------------------------------")
}

func (r *Rider) ShowEarnings() {
	fmt.Printf("💰 Rider: %s | Total Wallet: %.2f THB\n", r.Name, r.wallet)

	fmt.Println("--- Job History ---")
	for _, job := range r.History {
		// เรียก method GetStatus() หรือเรียก field ตรงๆ ก็ได้ (เพราะอยู่ package เดียวกัน)
		fmt.Printf("- Order %d: %s [%s]\n", job.Id, job.FoodName, job.status)
	}
	fmt.Println("-------------------")
}

func main() {
	// ลูกค้าสั่งอาหาร
	order1 := NewOrder(1, "Khao Man Gai", 55)
	order2 := NewOrder(2, "Tom Yum Kung", 90)

	// ไรเดอร์ภาคิน เข้าสู่ระบบ
	rider := NewRider("Phakin")

	// --- งานที่ 1 ---
	rider.AcceptJob(order1) // รับงาน (สถานะเปลี่ยน, เงินยังไม่เข้า)
	// สมมติว่าขับรถไปส่ง...
	rider.FinishJob(order1) // ส่งถึงที่ (สถานะ Delivered, เงินเข้า 44 บาท)

	// --- งานที่ 2 ---
	rider.AcceptJob(order2) // รับงาน
	rider.FinishJob(order2) // ส่งถึงที่ (เงินเข้า 72 บาท)

	// สรุปยอดเงิน
	rider.ShowEarnings()
}
