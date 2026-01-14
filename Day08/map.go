package main

import "fmt"

func main() {
	products := make(map[string]float64)

	products["iPhone"] = 15000
	products["Macbook"] = 45000
	products["ipad"] = 16000

	fmt.Println("List: ", products)

	fmt.Println("Price ipad:", products["ipad"])
	fmt.Println("price of macbook:", products["Macbook"])

	delete(products, "ipad")

	fmt.Println("product:", products)

	price, ok := products["samsung"]

	if ok {
		fmt.Printf("s price %.1f\n", price)
	} else {
		fmt.Println("invalid")
	}

	for product, price := range products {
		fmt.Printf("- %s: %.2f thb\n", product, price)
	}

}
