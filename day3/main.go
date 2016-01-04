package main

import (
	"fmt"

	"github.com/k3nn7/advent-of-code-go-2015/util"
)

type House struct {
	X int
	Y int
}

type Cursor struct {
	X int
	Y int
}

func (c *Cursor) Cmd(cmd string) {
	switch cmd {
	case "^":
		c.Y++
	case "v":
		c.Y--
	case ">":
		c.X++
	case "<":
		c.X--
	}
}

type World struct {
	visited []House
}

func (w *World) VisitedHouses() int {
	return len(w.visited)
}

func (w *World) Visit(h House) {
	if !w.wasVisited(h) {
		w.visited = append(w.visited, h)
	}
}

func (w *World) wasVisited(h House) bool {
	for _, v := range w.visited {
		if v == h {
			return true
		}
	}
	return false
}

func main() {
	data, err := util.ReadInput("./input.txt")
	if err != nil {
		panic(err)
	}

	part1Answer := part1VisitedHouses(data)
	part2Answer := part2VisitedHouses(data)

	fmt.Printf("Visited houses for part1: %d\n", part1Answer)
	fmt.Printf("Visited houses for part2: %d\n", part2Answer)
}

func part1VisitedHouses(data string) int {
	santa := Cursor{}
	world := World{}
	world.Visit(House{santa.X, santa.Y})
	for _, c := range data {
		santa.Cmd(string(c))
		world.Visit(House{santa.X, santa.Y})
	}
	return world.VisitedHouses()
}

func part2VisitedHouses(data string) int {
	santa := Cursor{}
	robot := Cursor{}
	world := World{}

	world.Visit(House{santa.X, santa.Y})
	world.Visit(House{robot.X, robot.Y})

	for i, c := range data {
		if i%2 == 0 {
			santa.Cmd(string(c))
			world.Visit(House{santa.X, santa.Y})
		} else {
			robot.Cmd(string(c))
			world.Visit(House{robot.X, robot.Y})
		}
	}
	return world.VisitedHouses()
}
