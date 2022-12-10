package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
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

func copyToClipboard(text string) {
	if runtime.GOOS != "linux" {
		panic("have you heard of our lord and saviour, GNU/Linux?")
	}

	if _, err := exec.LookPath("xsel"); err != nil {
		panic("paru -Syu xsel")
	}

	command := exec.Command("xsel", "-bi")
	stdin, err := command.StdinPipe()
	panicIf(err)

	err = command.Start()
	panicIf(err)

	io.WriteString(stdin, text)
	stdin.Close() // Command won't exit until stdin is closed.
	err = command.Wait()
	panicIf(err)
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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		enemyMove, err := NewMove(parts[0])
		panicIf(err)

		outcome, err := NewOutcome(parts[1])
		panicIf(err)

		myMove := resultMapping[outcome][enemyMove]
		score += NewScore(myMove, outcome)
	}

	fmt.Printf("My score: %d ", score)
	copyToClipboard(fmt.Sprint(score))
	fmt.Println("(written to clipboard)")
}
