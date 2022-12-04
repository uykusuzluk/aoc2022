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
