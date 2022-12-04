package main

import (
	"fmt"
)

type TopScores [3]int

func (t *TopScores) Check(x int) {
	if x > t[0] {
		t.Add(x, 0)
		// return true, 0
	}

	if x > t[1] {
		t.Add(x, 1)
		// return true, 1
	}

	if x > t[2] {
		t.Add(x, 2)
		// return true, 2
	}

	// return false, -1
}

func (t *TopScores) Add(x, pos int) {
	t.Print()
	switch pos {
	case 0:
		t[2] = t[1]
		t[1] = t[0]
	case 1:
		t[2] = t[1]
	}
	t[pos] = x
	t.Print()
}

func (t TopScores) Sum() int {
	sum := 0
	for i := 0; i < len(t); i++ {
		sum += t[i]
	}
	return sum
}

func (t TopScores) Print() {
	fmt.Printf("\nTop Three List:\n")
	fmt.Printf("1: %d\n", t[0])
	fmt.Printf("2: %d\n", t[1])
	fmt.Printf("3: %d\n", t[2])
}
