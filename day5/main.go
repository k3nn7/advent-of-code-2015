package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/k3nn7/advent-of-code-go-2015/util"
)

func isNicePart1(text string) bool {
	return hasVowels(text) && hasDoubledLetters(text) && !hasNaughtySubstring(text)
}

func isNicePart2(text string) bool {
	return hasDoubleDoubledLetters(text) && hasNiceTriplet(text)
}

func hasVowels(text string) bool {
	hasVowels := regexp.MustCompile(`(.*?[aeiou].*?){3}`)
	return hasVowels.MatchString(text)
}

func hasDoubledLetters(text string) bool {
	previous := rune(text[1])
	for _, current := range text {
		if current == previous {
			return true
		}
		previous = current
	}
	return false
}

func hasDoubleDoubledLetters(text string) bool {
	pairsCounts := make(map[string]int)
	pairs := getAllPairs(text)
	for _, pair := range pairs {
		pairsCounts[pair] = strings.Count(text, pair)
	}

	for _, count := range pairsCounts {
		if count >= 2 {
			return true
		}
	}
	return false
}

func hasNiceTriplet(text string) bool {
	triples := getAllTriples(text)
	for _, triple := range triples {
		if triple[0] == triple[2] {
			return true
		}
	}
	return false
}

func getAllPairs(text string) []string {
	pairs := make([]string, 0)
	for i := 0; i < len(text)-1; i++ {
		pair := fmt.Sprintf("%s%s", string(text[i]), string(text[i+1]))
		pairs = append(pairs, pair)
	}
	return pairs
}

func getAllTriples(text string) []string {
	triples := make([]string, 0)
	for i := 0; i < len(text)-2; i++ {
		triple := fmt.Sprintf("%s%s%s", string(text[i]), string(text[i+1]), string(text[i+2]))
		triples = append(triples, triple)
	}
	return triples
}

func hasNaughtySubstring(text string) bool {
	notAllowed := []string{"ab", "cd", "pq", "xy"}
	for _, str := range notAllowed {
		if strings.Contains(text, str) {
			return true
		}
	}
	return false
}

func main() {
	input, err := util.ReadInput("./input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(input, "\n")
	part1Nice := 0
	part2Nice := 0
	for _, line := range lines {
		if isNicePart1(line) {
			part1Nice++
		}
		if isNicePart2(line) {
			part2Nice++
		}
	}

	fmt.Printf("Nice strings for part1: %d\n", part1Nice)
	fmt.Printf("Nice strings for part2: %d\n", part2Nice)
}
