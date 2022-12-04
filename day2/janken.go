package main

import (
	"log"
)

func RoundPoints(opp, us string) int {
	oppShape, usShape := strToShape(opp), strToShape(us)
	if oppShape == Invalid || usShape == Invalid {
		log.Fatalln("unknown signs: " + opp + " " + oppShape.String() + " - " + us + " " + usShape.String())
	}

	return usShape.Points() + outcomePts(oppShape, usShape)
}

// X - Lose
// Y - Draw
// Z - Win
func RequiredShape(opp, result string) string {
	switch result {
	case "X":
		return strToShape(opp).shapeToLoseAgainst().shapeToStrUs()
	case "Y":
		return opp
	case "Z":
		return strToShape(opp).shapeToWinAgainst().shapeToStrUs()
	default:
		return ""
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
