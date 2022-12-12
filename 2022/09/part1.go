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

func (p Point) Step(d Direction) (newp Point) {
	newp = p
	switch d {
	case Up:
		newp.y += 1
	case Down:
		newp.y -= 1
	case Right:
		newp.x += 1
	case Left:
		newp.x -= 1
	}
	return
}

func (p Point) Follow(other Point) (newp Point) {
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

	newp = p

	dx := other.x - p.x
	dy := other.y - p.y

	if abs(dx) == 1 && abs(dy) == 2 {
		newp.x += dx
		newp.y += dec(dy)
	} else if abs(dx) == 2 && abs(dy) == 1 {
		newp.x += dec(dx)
		newp.y += dy
	} else if abs(dx) > 0 {
		newp.x += dec(dx)
	} else if abs(dy) > 0 {
		newp.y += dec(dy)
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

func PrettyPrint(xt, yt int, head Point, tail Point) string {
	formatted := ""
	// Print first line last, so that the diagram appears with (0,0) starting
	// on the bottom left.
	for y := yt - 1; y >= 0; y-- {
		for x := 0; x < xt; x++ {
			if head.x == x && head.y == y {
				formatted += fmt.Sprint("H")
			} else if tail.x == x && tail.y == y {
				formatted += fmt.Sprint("T")
			} else {
				formatted += fmt.Sprint(".")
			}
		}
		formatted += "\n"
	}
	return formatted
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	head, tail := Point{0, 0}, Point{0, 0}
	tailVisited := map[Point]bool{
		tail: true,
	}

	for scanner.Scan() {
		line := scanner.Text()
		move, err := parseMove(line)

		if err != nil {
			log.Println(err)
			continue
		}

		for i := 0; i < move.amount; i++ {
			head = head.Step(move.direction)
			tail = tail.Follow(head)
			fmt.Printf("M: %+v, H: %v, T: %v\n", move, head, tail)
			tailVisited[tail] = true

			fmt.Println(PrettyPrint(6, 5, head, tail))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("reading stdin failed: %v", err)
	}

	fmt.Printf("Tail visited %d fields.\n", len(tailVisited))
}
