package main

import (
	"fmt"
)

func main() {
	number := []int{10, 20, 30}
	fmt.Println("Number:", number)
	fmt.Println()

	score := make([]int, 3)
	fmt.Println("score:", score)
	fmt.Println()

	grades := make([]int, 3, 5)
	fmt.Printf("Grades: %v (len=%d, cap=%d)\n", grades, len(grades), cap(grades))
	fmt.Println()

	var emtySlice []int
	fmt.Println("Nil slice:", emtySlice)
	fmt.Println("Is nil?:", emtySlice == nil)
	fmt.Println()

	emtySlice = append(emtySlice, 10, 20, 30)
	fmt.Println(emtySlice)
	fmt.Println()

	fmt.Println("First number:", number[0])
	fmt.Println("Second number:", number[1])
	fmt.Println()

	number[2] = 25
	fmt.Println("Modified:", number)
	fmt.Println()

	fmt.Printf("Length: %d, Capacity: %d\n", len(number), cap(number))

	students := []string{"kin", "gulf", "gus", "team"}
	fmt.Println("Students name:", students)
	fmt.Println()

	students = append(students, "bas", "ice")
	fmt.Println("After append", students)
	fmt.Println()

	moreStudens := []string{"dew", "Aom"}
	students = append(students, moreStudens...)
	fmt.Println("After append slice:", students)
	fmt.Println()

	original := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", original)
	fmt.Println()

	sub1 := original[1:3]
	fmt.Println("Original[1:3]:", sub1)
	fmt.Println()

	sub2 := original[:4]
	fmt.Println("Original[:4]:", sub2)
	fmt.Println()

	sub3 := original[1:]
	fmt.Println("Original[1:]", sub3)
	fmt.Println()

	sub4 := original[:]
	fmt.Println("Original[:]:", sub4)
	fmt.Println()

	source := []int{1, 2, 3}
	destination := make([]int, len(source))

	copied := copy(destination, source)
	fmt.Printf("Copied %d elements\n", copied)
	fmt.Println("Source:", source)
	fmt.Println("Destination:", destination)
	fmt.Println()

	numbers2 := []int{10, 20, 30, 40, 50, 60}
	indexToDelete := 2
	numbers2 = append(numbers2[:indexToDelete], numbers2[indexToDelete+1:]...)
	fmt.Println("delete:", numbers2)
}
