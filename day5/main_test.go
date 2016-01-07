package main

import "testing"

var part1Inputs = []struct {
	text string
	nice bool
}{
	{
		text: "ugknbfddgicrmopn",
		nice: true,
	},
	{
		text: "aaa",
		nice: true,
	},
	{
		text: "jchzalrnumimnmhp",
		nice: false,
	},
	{
		text: "haegwjzuvuyypxyu",
		nice: false,
	},
	{
		text: "dvszwmarrgswjxmb",
		nice: false,
	},
}

func TestPart1Inputs(t *testing.T) {
	for _, input := range part1Inputs {
		nice := isNicePart1(input.text)
		if nice != input.nice {
			t.Errorf("'%s' nice should be: %v but is %v", input.text, input.nice, nice)
		}
	}
}

var part2Inputs = []struct {
	text string
	nice bool
}{
	{
		text: "qjhvhtzxzqqjkmpb",
		nice: true,
	},
	{
		text: "xxyxx",
		nice: true,
	},
	{
		text: "uurcxstgmygtbstg",
		nice: false,
	},
	{
		text: "ieodomkazucvgmuy",
		nice: false,
	},
}

func TestPart2Inputs(t *testing.T) {
	for _, input := range part2Inputs {
		nice := isNicePart2(input.text)
		if nice != input.nice {
			t.Errorf("'%s' nice should be: %v but is %v", input.text, input.nice, nice)
		}
	}
}
