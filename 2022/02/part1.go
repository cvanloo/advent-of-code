package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func modneg(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
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
		fallthrough
	case "X":
		m = Rock
	case "B":
		fallthrough
	case "Y":
		m = Paper
	case "C":
		fallthrough
	case "Z":
		m = Scissors
	default:
		err = fmt.Errorf("invalid move")
	}

	return
}

type Outcome int

const (
	Lost Outcome = 0
	Draw         = 3
	Won          = 6
)

type Score int

func NewScore(move Move, outcome Outcome) Score {
	return Score(int(move) + int(outcome))
}

type Opponent struct {
	moves   []Move
	current int
}

func NewOpponent(moves []Move) Opponent {
	return Opponent{moves, 0}
}

// (Opponent - Me) % 3 = Outcome, Score = Move + Outcome
// 1 - 2 = -1 % 3 = 2 Win (2+6=8)
// 2 - 1 = 1 % 3 = 1 Lose (1+0=1)
func (o *Opponent) Play(hand Move) (out Outcome) {
	ownMove := o.moves[o.current]
	o.current++

	result := modneg(int(ownMove-hand), 3)

	switch result {
	case 0:
		out = Draw
	case 1:
		out = Lost
	case 2:
		out = Won
	default:
		panic("not a valid outcome")
	}

	return
}

func CreateGame(input io.Reader) (Opponent, []Move, error) {
	var enemyMoves []Move
	var myMoves []Move
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		// parts[0] is opponent move
		// parts[1] is my move
		enemyMove, err := NewMove(parts[0])
		if err != nil {
			return Opponent{}, nil, fmt.Errorf("failed to parse %s as move: %v\n", parts[0], err)
		}

		myMove, err := NewMove(parts[1])
		if err != nil {
			return Opponent{}, nil, fmt.Errorf("failed to parse %s as move: %v\n", parts[1], err)
		}

		enemyMoves = append(enemyMoves, enemyMove)
		myMoves = append(myMoves, myMove)
	}

	if err := scanner.Err(); err != nil {
		return Opponent{}, nil, err
	}

	return NewOpponent(enemyMoves), myMoves, nil
}

// A X Rock
// B Y Paper
// C Z Scissors
//
// A Y
// B X
// C Z
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

	opponent, moves, err := CreateGame(fd)
	if err != nil {
		log.Fatalf("failed to create game: %v\n", err)
	}

	score := Score(0)
	for _, move := range moves {
		outcome := opponent.Play(move)
		score += NewScore(move, outcome)
	}

	fmt.Printf("My score: %d\n", score)
}
