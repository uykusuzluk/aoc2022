package main

import "github.com/uykusuzluk/aoc2022/aocutils"

func main() {
	input := "day2_input"
	linesCh := make(chan string, 50)
	go aocutils.InputLinesFromFile(input, linesCh)
	// totalScore(linesCh)
	totalPriorities(linesCh)
}

func totalPriorities(q <-chan string) {
	for line := range q {
		// FIXME: Not utf-8 safe!
		firstComp, secondComp := line[:len(line)/2], line[len(line)/2:]

	}
}
