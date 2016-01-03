package main

import "testing"

func TestSantaStartsOnGroundFloor(t *testing.T) {
	santa := Santa{}
	if santa.Floor != 0 {
		t.Errorf("Santa should start at floor 0 but got: %d", santa.Floor)
	}
}

var useCases = []struct {
	input string
	floor int
}{
	{
		input: "(())",
		floor: 0,
	},
	{
		input: "(((",
		floor: 3,
	},
	{
		input: "(()(()(",
		floor: 3,
	},
	{
		input: "())",
		floor: -1,
	},
	{
		input: ")())())",
		floor: -3,
	},
}

func TestSantaParseInput(t *testing.T) {
	for _, useCase := range useCases {
		santa := Santa{}
		santa.ParseInput(useCase.input)
		if useCase.floor != santa.Floor {
			t.Errorf("Santa should be on floor %d but is on %d", useCase.floor, santa.Floor)
		}
	}
}

var whenEnterMinusOneCases = []struct {
	input string
	when  int
}{
	{
		input: ")",
		when:  1,
	},
	{
		input: "()())",
		when:  5,
	},
	{
		input: ")()",
		when:  1,
	},
}

func TestWhenEnteredMinusOne(t *testing.T) {
	for _, useCase := range whenEnterMinusOneCases {
		santa := Santa{}
		santa.ParseInput(useCase.input)
		if useCase.when != santa.WhenEnteredMinusOne {
			t.Errorf("Santa should enter -1 floor after %d move but he entered after %d", useCase.when, santa.WhenEnteredMinusOne)
		}
	}
}
