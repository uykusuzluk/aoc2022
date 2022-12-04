package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/uykusuzluk/aoc2022/aocutils"
)

func main() {
	input := "day2_input"
	linesCh := make(chan string, 50)
	go aocutils.InputLinesFromFile(input, linesCh)
	// totalScore(linesCh)
	followGuide(linesCh)
}

func totalScore(q <-chan string) {
	var pts int
	for line := range q {
		signs := strings.Split(line, " ")
		if len(signs) <= 1 {
			log.Println("error splitting line: " + line)
		}
		roundPts := RoundPoints(signs[0], signs[1])
		log.Println("This round: " + strconv.Itoa(roundPts))
		pts += roundPts
	}
	fmt.Println("Final score: " + strconv.Itoa(pts))
}

func followGuide(q <-chan string) {
	var pts int
	for line := range q {
		signs := strings.Split(line, " ")
		if len(signs) <= 1 {
			log.Println("error splitting line: " + line)
		}
		shp := RequiredShape(signs[0], signs[1])
		roundPts := RoundPoints(signs[0], shp)
		log.Println("This round: " + strconv.Itoa(roundPts))
		pts += roundPts
	}
	fmt.Println("Final score: " + strconv.Itoa(pts))
}
