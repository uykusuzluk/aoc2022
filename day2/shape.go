package main

type Shape int

const (
	Invalid Shape = iota
	Rock
	Paper
	Scissors
)

func (s Shape) String() string {
	switch s {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	default:
		return "Invalid"
	}
}

func (s Shape) shapeToStrUs() string {
	switch s {
	case Rock:
		return "X"
	case Paper:
		return "Y"
	case Scissors:
		return "Z"
	default:
		return ""
	}
}

func strToShape(str string) Shape {
	switch str {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		return Invalid
	}
}

func (s Shape) Points() int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		return 0
	}
}

func (s Shape) shapeToWinAgainst() Shape {
	switch s {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		return Invalid
	}
}

func (s Shape) shapeToLoseAgainst() Shape {
	switch s {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		return Invalid
	}
}
