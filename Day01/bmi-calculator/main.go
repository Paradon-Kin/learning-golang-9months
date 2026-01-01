package main

import "fmt"

func main() {
	var weight float64
	var height float64

	fmt.Println("กรุณาใส่น้ำหนัก (kg): ")
	fmt.Scan(&weight)

	fmt.Println("กรุณาใส่ส่วนสูง (m): ")
	fmt.Scan(&height)

	bmi := weight / (height * height)

	fmt.Printf("BMI ของคุณคือ: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("คุณผอมเกินไป")
	} else if bmi < 25 {
		fmt.Println("คุณน้ำหนักปกติ")
	} else if bmi < 30 {
		fmt.Println("คุณน้ำหนักเกิน")
	} else {
		fmt.Println("คุณอ้วน")
	}
}
