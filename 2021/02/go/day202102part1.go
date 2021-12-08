package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	depth, horiz int
}

type direction int

const (
	forward direction = iota
	up
	down
)

type command struct {
	direction direction
	amount    int
}

func (p *position) handle(cmd command) {
	switch cmd.direction {
	case forward:
		p.horiz += cmd.amount
	case up:
		p.depth -= cmd.amount
	case down:
		p.depth += cmd.amount
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	panicIf(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pos := position{}

	for scanner.Scan() {
		line := scanner.Text()
		instStrArr := strings.Fields(line)

		var d direction

		switch instStrArr[0] {
		case "forward":
			d = forward
		case "up":
			d = up
		case "down":
			d = down
		}

		u, err := strconv.Atoi(instStrArr[1])
		panicIf(err)

		inst := command{
			direction: d,
			amount:    u,
		}

		pos.handle(inst)
	}

	fmt.Println("Result:", pos.depth*pos.horiz)
}
