package main

import (
	"fmt"
)

func main() {
	var number [5]int
	fmt.Println("Emty array:", number)
	fmt.Println()

	score := [3]int{85, 90, 78}
	fmt.Println("Score:", score)
	fmt.Println()

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	fmt.Println("Days:", days)
	fmt.Println()

	fmt.Println("Fisrt score:", score[0])
	fmt.Println("Second score:", score[1])
	fmt.Println()

	score[2] = 70
	fmt.Println("Update score:", score[2])
	fmt.Println()

	fmt.Println("Length of score:", len(score))
	fmt.Println("Length of Days:", len(days))
	fmt.Println()

	fmt.Println("Score using for loop")
	for i := 0; i < len(score); i++ {
		fmt.Printf("Score[%d] = %d\n", i, score[i])
	}
	fmt.Println()

	fmt.Println("Days using range:")
	for index, day := range days {
		fmt.Printf("%d: %s\n", index, day)
	}
	fmt.Println()

	for index, score := range score {
		fmt.Printf("%d: %d\n", index, score)
	}
	fmt.Println()
}
