package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/k3nn7/advent-of-code-go-2015/util"
)

type Dimensions struct {
	L int
	W int
	H int
}

func (d *Dimensions) PaperRequired() int {
	sideA := d.L * d.W
	sideB := d.W * d.H
	sideC := d.H * d.L
	slack := math.Min(math.Min(float64(sideA), float64(sideB)), float64(sideC))
	return 2*sideA + 2*sideB + 2*sideC + int(slack)
}

func (d *Dimensions) RibbonRequired() int {
	sides := []int{d.L, d.W, d.H}
	sort.Ints(sides)
	bow := d.L * d.W * d.H
	return 2*sides[0] + 2*sides[1] + bow
}

func ParseDimensions(input string) (dimension Dimensions) {
	splitted := strings.Split(input, "x")
	l, _ := strconv.Atoi(splitted[0])
	w, _ := strconv.Atoi(splitted[1])
	h, _ := strconv.Atoi(splitted[2])

	return Dimensions{l, w, h}
}

func main() {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(input, "\n")

	totalPaperRequired := 0
	totalRibbonRequired := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		d := ParseDimensions(line)
		totalPaperRequired += d.PaperRequired()
		totalRibbonRequired += d.RibbonRequired()
	}
	fmt.Printf("Total paper required: %d\n", totalPaperRequired)
	fmt.Printf("Total ribbon required: %d\n", totalRibbonRequired)
}
