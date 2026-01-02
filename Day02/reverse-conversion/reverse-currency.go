package main

import "fmt"

// Exchange rate constants (as of January 2025)
const (
	THBtoUSD = 0.029
	THBtoJPY = 4.32
	THBtoEUR = 0.027
)

const FeePercentage = 0.01

func main() {
	var thbAmount float64

	fmt.Println("======================================")
	fmt.Println("    Currency Converter (with Fee)     ")
	fmt.Println("  เครื่องแปลงสกุลเงิน (มีค่าธรรมเนียม) ")
	fmt.Println("======================================")
	fmt.Printf("ค่าธรรมเนียม / Fee Rate: %.0f%%\n", FeePercentage*100)
	fmt.Println()

	// Get THB amount from user
	fmt.Println("ใส่จำนวนเงิน (THB):")
	fmt.Scan(&thbAmount)

	fee := thbAmount * FeePercentage

	// Validate: check for negative or zero amount
	if thbAmount <= 0 {
		fmt.Println("❌ กรุณาใส่จำนวนเงินที่มากกว่า 0")
		return
	}

	// Validate: check if user has enough money (including fee)
	if thbAmount < fee {
		fmt.Println("❌ เงินไม่พอแลก")
		fmt.Printf("ต้องการขั้นต่ำ %.2f บาท (รวมค่าธรรมเนียม)\n", fee)
		return
	}

	// Calculate amount after deducting fee
	thbAfterFee := thbAmount - fee
	fmt.Printf("จำนวนเงินหลังจากหักค่าธรรมเนียม: %.2f THB\n", thbAfterFee)

	usdAmount := thbAfterFee * THBtoUSD
	jpyAmount := thbAfterFee * THBtoJPY
	eurAmount := thbAfterFee * THBtoEUR

	// Display conversion results
	fmt.Println("\n======== Conversion Results ========")
	fmt.Println("========   ผลการแปลง   ========")
	fmt.Printf("USD: $%.2f\n", usdAmount)
	fmt.Printf("JPY: ¥%.2f\n", jpyAmount)
	fmt.Printf("EUR: €%.2f\n", eurAmount)

	// Display current exchange rates
	fmt.Println("\nCurrent Exchange Rates:")
	fmt.Printf("1 THB = %.3f USD\n", THBtoUSD)
	fmt.Printf("1 THB = %.2f JPY\n", THBtoJPY)
	fmt.Printf("1 THB = %.3f EUR\n", THBtoEUR)

}
