package main

import (
	"fmt"
	"time"
)

func runner(name string, speed time.Duration, ch chan string) {
	time.Sleep(speed)
	ch <- name + "เข้าเส้นชัย"
}

func main() {
	turtleChan := make(chan string)
	rabbitChan := make(chan string)

	go runner("turtle", 2*time.Second, turtleChan)
	go runner("rabbit", 1*time.Second, rabbitChan)

	fmt.Println("ออกตัว...")

	select {
	case msg1 := <-turtleChan:
		fmt.Println("winner is:", msg1)
	case msg2 := <-rabbitChan:
		fmt.Println("winner is:", msg2)

	}
}
