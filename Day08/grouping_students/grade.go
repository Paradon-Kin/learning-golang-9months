package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	students := make(map[string][]string)

	for {
		name := readLine("Name: ")
		if name == "exit" {
			break
		}
		grade := readLine("Grade: ")

		students[grade] = append(students[grade], name)
	}

	for grade, name := range students {
		fmt.Printf("%s : %s\n", grade, name)
	}
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
