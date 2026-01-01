package main

import (
	"fmt"
)

func main() {
	var height float64

	fmt.Println("กรุณาใส่ส่วนสูง (m): ")
	fmt.Scan(&height)

	//แปลงเป็น cm
	//height := heightcm / 100

	weight := 22 * (height * height)

	fmt.Printf("น้ำหนักที่เหมาะสม: %.1f kg\n", weight)

}
