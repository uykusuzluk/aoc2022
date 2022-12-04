package main

import (
	"log"
)

func RoundPoints(opp, us string) int {
	oppShape, usShape := strToShape(opp), strToShape(us)
	if oppShape == Invalid || usShape == Invalid {
		log.Fatalln("unknown signs: " + opp + " " + oppShape.String() + " - " + us + " " + usShape.String())
	}

	return shapePts(usShape) + outcomePts(oppShape, usShape)
}

// X - Lose
// Y - Draw
// Z - Win
func RequiredShape(opp, result string) string {
	switch result {
	case "X":
		return shapeToStrUs(shapeToLoseAgainst(strToShape(opp)))
	case "Y":
		return opp
	case "Z":
		return shapeToStrUs(shapeToWinAgainst(strToShape(opp)))
	default:
		return ""
	}
}

func shapeToWinAgainst(s Shape) Shape {
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

func shapeToLoseAgainst(s Shape) Shape {
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

func shapeToStrUs(s Shape) string {
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

func shapePts(sign Shape) int {
	switch sign {
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

func outcomePts(opp, us Shape) int {
	// Draw case
	if opp == us {
		return 3
	}

	switch opp {
	case Rock:
		if us == Paper {
			return 6
		}
		return 0
	case Paper:
		if us == Scissors {
			return 6
		}
		return 0
	case Scissors:
		if us == Rock {
			return 6
		}
		return 0
	default:
		return -1
	}
}
