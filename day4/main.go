package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func findHash(key, prefix string) int {
	for i := 0; ; i++ {
		input := fmt.Sprintf("%s%d", key, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if strings.HasPrefix(hash, prefix) {
			return i
		}
	}
}

func main() {
	input := "iwrupvqb"
	part1Answer := findHash(input, "00000")
	part2Answer := findHash(input, "000000")
	fmt.Printf("Part 1 answer: %d\n", part1Answer)
	fmt.Printf("Part 2 answer: %d\n", part2Answer)
}
