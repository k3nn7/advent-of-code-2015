package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/k3nn7/advent-of-code-go-2015/day6/binarygrid"
	"github.com/k3nn7/advent-of-code-go-2015/day6/brightnessgrid"
	"github.com/k3nn7/advent-of-code-go-2015/util"
)

type Command struct {
	Cmd string
	X1  int
	Y1  int
	X2  int
	Y2  int
}

func ParseCommand(cmd string) Command {
	re := regexp.MustCompile(`([a-z\s]+)(\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(cmd, -1)
	c := strings.Trim(matches[0][1], " ")
	x1, _ := strconv.Atoi(matches[0][2])
	y1, _ := strconv.Atoi(matches[0][3])
	x2, _ := strconv.Atoi(matches[0][4])
	y2, _ := strconv.Atoi(matches[0][5])
	return Command{
		Cmd: c,
		X1:  x1,
		Y1:  y1,
		X2:  x2,
		Y2:  y2,
	}
}

func passCommandsToGrid(grid LightGrid, commands []string) {
	for _, line := range commands {
		if len(line) < 5 {
			continue
		}
		c := ParseCommand(line)
		switch c.Cmd {
		case "turn on":
			grid.TurnOnSquare(c.X1, c.Y1, c.X2, c.Y2)
		case "turn off":
			grid.TurnOffSquare(c.X1, c.Y1, c.X2, c.Y2)
		case "toggle":
			grid.ToggleSquare(c.X1, c.Y1, c.X2, c.Y2)
		}
	}
}

func main() {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		panic(err)
	}
	commands := strings.Split(input, "\n")

	binaryGrid := binarygrid.NewGrid(1000)
	brightnessGrid := brightnessgrid.NewGrid(1000)

	passCommandsToGrid(binaryGrid, commands)
	passCommandsToGrid(brightnessGrid, commands)

	fmt.Printf("Total brightness (numbero of lit lights) for part1: %d\n", binaryGrid.TotalBrightness())
	fmt.Printf("Total brightness for part2: %d\n", brightnessGrid.TotalBrightness())
}
