package main

import (
	"fmt"
	"time"
)

func makeCoffee(menu string, ch chan string) {
	fmt.Printf("The Barista is making %s ...\n", menu)
	time.Sleep(3 * time.Second)
	ch <- menu + " is Ready!"
}

func main() {
	oderChan := make(chan string)

	go makeCoffee("Latte", oderChan)
	go makeCoffee("Espresso", oderChan)
	go makeCoffee("Mocha", oderChan)

	fmt.Println("Waiting for drink...")

	menu1 := <-oderChan
	fmt.Println("Receive:", menu1)

	menu2 := <-oderChan
	fmt.Println("Receive:", menu2)

	menu3 := <-oderChan
	fmt.Println("Receive:", menu3)

	fmt.Println("All full!")
}
