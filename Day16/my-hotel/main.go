package main // <--- ⚠️ ต้องเป็น main เท่านั้น ห้ามเป็น myhotel

import (
	room "hotel/hotel" // <--- เรียกชื่อ module/ชื่อโฟลเดอร์
)

func main() {
	room1 := room.NewRoom(101, 2500)

	room1.CheckIn("Kin")
	room1.CheckOut(3)

}
