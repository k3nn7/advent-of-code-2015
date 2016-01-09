package main

import (
	"reflect"
	"testing"
)

var useCases = []struct {
	command  string
	expected Command
}{
	{
		command:  "turn off 370,39 through 425,839",
		expected: Command{"turn off", 370, 39, 425, 839},
	},
	{
		command:  "toggle 467,662 through 555,957",
		expected: Command{"toggle", 467, 662, 555, 957},
	},
	{
		command:  "turn on 975,2 through 984,623",
		expected: Command{"turn on", 975, 2, 984, 623},
	},
}

func TestParseCommand(t *testing.T) {
	for _, useCase := range useCases {
		command := ParseCommand(useCase.command)
		if !reflect.DeepEqual(useCase.expected, command) {
			t.Errorf("Expected %#v got %#v", useCase.expected, command)
		}
	}
}
