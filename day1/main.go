package main

import (
	"fmt"

	"github.com/k3nn7/advent-of-code-go-2015/util"
)

type Santa struct {
	Floor               int
	WhenEnteredMinusOne int
}

func (s *Santa) ParseInput(input string) {
	for i, char := range input {
		switch string(char) {
		case "(":
			s.Floor++
		case ")":
			s.Floor--
		}

		if s.Floor == -1 && s.WhenEnteredMinusOne == 0 {
			s.WhenEnteredMinusOne = i + 1
		}
	}
}

func main() {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		panic(err)
	}
	santa := Santa{}
	santa.ParseInput(input)
	fmt.Printf("Santa is on: %d floor\n", santa.Floor)
	fmt.Printf("Santa entered -1 floor after: %d moves\n", santa.WhenEnteredMinusOne)
}
