package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSquare(number int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	result := number * number
	ch <- fmt.Sprintf("Number %d squared = %d", number, result)
}

func main() {
	ch := make(chan string, 10)

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go calculateSquare(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
