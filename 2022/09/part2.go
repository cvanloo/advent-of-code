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

func (p Point) Distance(to Point) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(to.x-p.x) + abs(to.y-p.y)
}

func (p *Point) Step(d Direction) {
	switch d {
	case Up:
		p.y += 1
	case Down:
		p.y -= 1
	case Right:
		p.x += 1
	case Left:
		p.x -= 1
	}
	return
}

func (p *Point) Follow(other Point) {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	dec := func(n int) int {
		if n < 0 {
			return n + 1
		}
		return n - 1
	}

	dx := other.x - p.x
	dy := other.y - p.y

	if abs(dx) == 1 && abs(dy) == 2 {
		p.x += dx
		p.y += dec(dy)
	} else if abs(dx) == 2 && abs(dy) == 1 {
		p.x += dec(dx)
		p.y += dy
	} else if abs(dx) == 2 && abs(dy) == 2 {
		p.x += 1
		p.y += 1
	} else if abs(dx) > 0 {
		p.x += dec(dx)
	} else if abs(dy) > 0 {
		p.y += dec(dy)
	}
	return
}

func (p Point) String() string {
	return fmt.Sprintf("(%d %d)", p.x, p.y)
}

type Direction string

const (
	Up    Direction = "Up"
	Down            = "Down"
	Left            = "Left"
	Right           = "Right"
)

type Move struct {
	direction Direction
	amount    int
}

func parseMove(line string) (m Move, err error) {
	fields := strings.Fields(line)

	if l := len(fields); l < 2 {
		err = fmt.Errorf("expected two fields, got: %d", l)
		return
	}

	direction := fields[0]
	amount, err := strconv.Atoi(fields[1])

	if err != nil {
		err = fmt.Errorf("move amount must be a number: %v", err)
		return
	}

	switch direction {
	case "U":
		m = Move{Up, amount}
	case "D":
		m = Move{Down, amount}
	case "L":
		m = Move{Left, amount}
	case "R":
		m = Move{Right, amount}
	default:
		err = fmt.Errorf("move direction must be one of UDLR, got: %s", direction)
		return
	}
	return
}

func PrettyPrintField(knots [10]Point) string {
	formatted := ""

	// y 5, x 6
	for y := 21; y >= 0; y-- {
		for x := 0; x < 26; x++ {
			for i, knot := range knots {
				if knot.x == x && knot.y == y {
					formatted += fmt.Sprint(i)
					goto cont
				}
			}
			formatted += "."
		cont:
		}
		formatted += "\n"
	}

	return formatted
}

func PrettyPrint(positionsVisited map[Point]bool) string {
	formatted := ""

	visited := []Point{}
	lx := 0
	hx := 0
	ly := 0
	hy := 0

	for k := range positionsVisited {
		visited = append(visited, k)
		if k.x < lx {
			lx = k.x
		} else if k.x > hx {
			hx = k.x
		}
		if k.y < ly {
			ly = k.y
		} else if k.y > hy {
			hy = k.y
		}
	}

	for y := hy; y >= ly; y-- {
		for x := lx; x <= hx; x++ {
			if positionsVisited[Point{x, y}] {
				formatted += "#"
			} else {
				formatted += "."
			}
		}
		formatted += "\n"
	}
	return formatted
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	knots := [10]Point{}
	head := &knots[0]
	tail := &knots[9]
	positionsVisited := map[Point]bool{
		*head: true,
	}

	for scanner.Scan() {
		line := scanner.Text()
		move, err := parseMove(line)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("M: %+v, H: %v, T: %v\n", move, head, tail)
		for i := 0; i < move.amount; i++ {
			head.Step(move.direction)
			for j := 0; j < 10-1; j++ {
				followee := &knots[j]
				follower := &knots[j+1]

				follower.Follow(*followee)
			}
			positionsVisited[*tail] = true
		}
		fmt.Println(PrettyPrintField(knots))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("reading stdin failed: %v", err)
	}

	fmt.Print(PrettyPrint(positionsVisited))

	fmt.Printf("Tail visited %d fields.\n", len(positionsVisited))
}
