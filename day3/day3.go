package main

import (
	"fmt"
	"strconv"

	"github.com/uykusuzluk/aoc2022/aocutils"
)

func main() {
	input := "day3input"
	linesCh := make(chan string, 50)
	go aocutils.InputLinesFromFile(input, linesCh)
	totalPriorities(linesCh)
}

func totalPriorities(q <-chan string) {
	var acc int
	for line := range q {
		acc += Rucksack(line).Priority()
	}

	fmt.Println("Total priorities: " + strconv.Itoa(acc))
}
