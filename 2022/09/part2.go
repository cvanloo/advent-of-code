package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p *Point) Add(other Point) {
	p.x += other.x
	p.y += other.y
}

var directionMap = map[string]Point{
	"U": {0, 1},
	"D": {0, -1},
	"L": {-1, 0},
	"R": {1, 0},
}

type Move struct {
	by     Point
	repeat int
}

func parseMove(line string) (m Move, err error) {
	fields := strings.Fields(line)
	if l := len(fields); l < 2 {
		return m, fmt.Errorf("expected two fields, got: %d", l)
	}

	direction := fields[0]
	amount, err := strconv.Atoi(fields[1])
	if err != nil {
		return m, fmt.Errorf("move amount must be a number: %v", err)
	}

	byPoint, ok := directionMap[direction]
	if !ok {
		return m, fmt.Errorf("direction must be one of UDLR, got: %s", direction)
	}
	return Move{byPoint, amount}, nil
}

func followPoint(p *Point, other Point) {
	// Go's std doesn't have an abs for ints :-(
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	dx := other.x - p.x
	dy := other.y - p.y

	if abs(dx) < 2 && abs(dy) < 2 {
		return
	}

	if other.x > p.x {
		p.x += 1
	}
	if other.x < p.x {
		p.x -= 1
	}

	if other.y > p.y {
		p.y += 1
	}
	if other.y < p.y {
		p.y -= 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	knots := [10]Point{}
	head := &knots[0]
	tail := &knots[9]
	positionsVisited := map[Point]bool{*head: true}

	for scanner.Scan() {
		line := scanner.Text()
		move, err := parseMove(line)

		if err != nil {
			log.Println(err)
			continue
		}

		for i := 0; i < move.repeat; i++ {
			head.Add(move.by)
			for j := 0; j < 10-1; j++ {
				followee := &knots[j]
				follower := &knots[j+1]
				followPoint(follower, *followee)
			}
			positionsVisited[*tail] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("reading stdin failed: %v", err)
	}

	fmt.Printf("Tail visited %d fields.\n", len(positionsVisited))
}
