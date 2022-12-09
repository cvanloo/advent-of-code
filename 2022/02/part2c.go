package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type Move int

const (
	Rock Move = iota + 1
	Paper
	Scissors
)

func NewMove(str string) (m Move, err error) {
	switch str {
	case "A":
		m = Rock
	case "B":
		m = Paper
	case "C":
		m = Scissors
	default:
		err = fmt.Errorf("invalid move: %s", str)
	}

	return
}

type Outcome int

const (
	Lost Outcome = 0
	Draw         = 3
	Won          = 6
)

func NewOutcome(str string) (Outcome, error) {
	switch str {
	case "X":
		return Lost, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Won, nil
	}
	return Outcome(0), fmt.Errorf("invalid outcome: %s", str)
}

type Score int

func NewScore(move Move, outcome Outcome) Score {
	return Score(int(move) + int(outcome))
}

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	resultMapping := map[Outcome]map[Move]Move{
		Lost: {
			Rock:     Scissors,
			Paper:    Rock,
			Scissors: Paper,
		},
		Draw: {
			Rock:     Rock,
			Paper:    Paper,
			Scissors: Scissors,
		},
		Won: {
			Rock:     Paper,
			Paper:    Scissors,
			Scissors: Rock,
		},
	}

	score := Score(0)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		parts := strings.Split(line, " ")

		enemyMove, err := NewMove(parts[0])
		panicIf(err)

		outcome, err := NewOutcome(parts[1])
		panicIf(err)

		myMove := resultMapping[outcome][enemyMove]
		score += NewScore(myMove, outcome)
	}

	fmt.Printf("My score: %d\n", score)
}
