package main

import (
	"fmt"
	"sync"
	"time"
)

func dowloadFile(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start downloading [%s]...\n", name)
	time.Sleep(2 * time.Second)
	fmt.Printf("[%s]Dowloaded!\n", name)
}

func main() {
	var wg sync.WaitGroup

	file := []string{"Movie.mp4", "Music.mp3", "Image.jpg"}

	for _, f := range file {
		wg.Add(1)
		go dowloadFile(f, &wg)
	}

	fmt.Println("⏳ Waiting for downloads...")
	wg.Wait()

	fmt.Println("🎉 All downloads finished!")
}
