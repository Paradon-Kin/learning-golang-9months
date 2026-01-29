package main

import "fmt"

type student struct {
	name  string
	score int
}

func main() {
	var s = []student{
		{name: "kin", score: 50},
		{name: "gus", score: 45},
	}
	for i := range s {
		bonus(&s[i])
		fmt.Println("", s)
	}

}

func bonus(s *student) {
	s.score += 10
}
