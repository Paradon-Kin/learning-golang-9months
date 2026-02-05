package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fmt.Sprintf("Result from worker %d", id)
}

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Receive", msg)
	}
}
