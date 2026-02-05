package main

import (
	"fmt"
	"time"
)

func makePizza(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Pizza Hawaiian"
}

func main() {
	ch := make(chan string)

	fmt.Println("Ordering Pizza...")

	go makePizza(ch)

	select {
	case pizzza := <-ch:
		fmt.Println("Yummy! Receive:", pizzza)
	case <-time.After(4 * time.Second):
		fmt.Println("Timeout! I'm starving")
	}
}
