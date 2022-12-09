package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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

func CreateGame(input io.Reader) ([]Move, []Outcome, error) {
	var enemyMoves []Move
	var outcomes []Outcome
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		// parts[0] is opponent move
		// parts[1] is my move
		enemyMove, err := NewMove(parts[0])
		if err != nil {
			return []Move{}, nil, fmt.Errorf("failed to parse %s as move: %v\n", parts[0], err)
		}

		outcome, err := NewOutcome(parts[1])
		if err != nil {
			return []Move{}, nil, fmt.Errorf("failed to parse %s as outcome: %v\n", parts[1], err)
		}

		enemyMoves = append(enemyMoves, enemyMove)
		outcomes = append(outcomes, outcome)
	}

	if err := scanner.Err(); err != nil {
		return []Move{}, nil, err
	}

	return enemyMoves, outcomes, nil
}

// A Rock
// B Paper
// C Scissors
//
// X Lose
// Y Draw
// Z Win
//
// Rock 1, Paper 2, Scissors 3
// Lose 0, Draw 3, Win 6
func main() {
	fd, err := os.Open("input.txt")
	//fd, err := os.Open("test.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v\n", err)
	}
	defer fd.Close()

	enemyMoves, outcomes, err := CreateGame(fd)
	if err != nil {
		log.Fatalf("failed to create game: %v\n", err)
	}

	// I don't know how to solve...
	// (enemyMove - myMove) % 3 = Outcome
	// ... for myMove
	// so instead I just use a map.
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
	for i, enemyMove := range enemyMoves {
		outcome := outcomes[i]
		myMove := resultMapping[outcome][enemyMove]
		score += NewScore(myMove, outcome)
	}

	fmt.Printf("My score: %d\n", score)
}
