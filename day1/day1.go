package main

import (
	"log"
	"strconv"

	"github.com/uykusuzluk/aoc2022/aocutils"
)

func main() {
	// input := "https://adventofcode.com/2022/day/1/input"
	input := "day1input"
	linesCh := make(chan string, 50)
	go aocutils.InputLinesFromFile(input, linesCh)
	// maxCalories(linesCh)
	maxThreeCalories(linesCh)
}

func maxCalories(q <-chan string) {
	max, acc := 0, 0
	for line := range q {
		if line == "" {
			if max < acc {
				max = acc
			}
			acc = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Println("cannot convert str to int. line: " + line)
			continue
		}

		acc += calories
	}

	log.Println("max calories: " + strconv.Itoa(max))
}

func maxThreeCalories(q <-chan string) {
	var (
		acc int
		t   TopScores
	)
	for line := range q {
		if line == "" {
			t.Check(acc)
			acc = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Println("cannot convert str to int. line: " + line)
			continue
		}

		acc += calories
	}

	log.Println("max calories: " + strconv.Itoa(t.Sum()))
}
