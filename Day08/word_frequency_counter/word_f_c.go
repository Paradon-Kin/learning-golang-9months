package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	voteList := make(map[string]int)

	for {
		vote := readLine("vote: ")
		if vote == "exit" {
			break
		}
		voteList[vote]++
	}

	var winner string
	var maxScore = 0

	for n, s := range voteList {
		fmt.Printf("%s : %d\n", n, s)
		if s > maxScore {
			maxScore = s
			winner = n
		}
	}

	if maxScore == 0 {
		fmt.Println("No one won")
	} else {
		fmt.Printf("Winner is: %s\n", winner)
	}

}
func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
