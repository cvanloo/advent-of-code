package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Destructure() (int, int) {
	return p.x, p.y
}

func (p Point) Equal(other Point) bool {
	return p.x == other.x && p.y == other.y
}

type Structure uint

const (
	Air Structure = iota
	Wall
	Sand
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type strconvWrapper struct {
	err error
}

func (str *strconvWrapper) Atoi(s string) int {
	if str.err == nil {
		i, err := strconv.Atoi(s)
		str.err = err
		return i
	}
	return 0
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(input string) (fields map[Point]Structure, lowestPoint int, err error) {
	fields = make(map[Point]Structure)
	intParser := strconvWrapper{nil}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "\n" {
			continue
		}

		points := strings.Split(line, " -> ")
		for i := 1; i < len(points); i++ {
			p1s := points[i-1]
			p2s := points[i]

			p1 := strings.Split(p1s, ",")
			p2 := strings.Split(p2s, ",")

			startX := intParser.Atoi(p1[0])
			startY := intParser.Atoi(p1[1])
			endX := intParser.Atoi(p2[0])
			endY := intParser.Atoi(p2[1])
			if err = intParser.err; err != nil {
				return
			}

			if startX > endX {
				startX, endX = endX, startX
			}
			if startY > endY {
				startY, endY = endY, startY
			}

			if startX == endX {
				for i := startY; i <= endY; i++ {
					fields[Point{startX, i}] = Wall
				}
			} else if startY == endY {
				for i := startX; i <= endX; i++ {
					fields[Point{i, startY}] = Wall
				}
			} else {
				panic("diagonals are not allowed")
			}

			if startY > lowestPoint {
				lowestPoint = startY
			}
			if endY > lowestPoint {
				lowestPoint = endY
			}
		}
	}
	return
}

func simulateSand(fields map[Point]Structure, spawnPoint Point, lowestPoint int) uint {
	var sandAtRest uint

	for {
		x, y := spawnPoint.Destructure()
		for {
			if y == lowestPoint-1 {
				break
			}

			if fields[Point{x, y + 1}] == Air {
				y += 1
			} else if fields[Point{x - 1, y + 1}] == Air {
				x -= 1
				y += 1
			} else if fields[Point{x + 1, y + 1}] == Air {
				x += 1
				y += 1
			} else {
				break
			}
		}

		fields[Point{x, y}] = Sand
		sandAtRest += 1
		//printField(fields, lowestPoint)

		if spawnPoint.Equal(Point{x, y}) {
			break
		}
	}

	return sandAtRest
}

func printField(fields map[Point]Structure, lowestPoint int) {
	for y := 0; y < 12; y++ {
		for x := 480; x < 520; x++ {
			if x == 500 && y == 0 {
				fmt.Print("+")
				continue
			}
			if y == lowestPoint {
				fmt.Print("█")
				continue
			}
			switch (fields[Point{x, y}]) {
			case Air:
				fmt.Print(".")
			case Wall:
				fmt.Print("#")
			case Sand:
				fmt.Print("o")
			default:
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	panicIf(err)
	fields, lowestPoint, err := parseInput(string(input))
	panicIf(err)
	fmt.Printf("Number of wall fields parsed: %d\n", len(fields))
	result := simulateSand(fields, Point{500, 0}, lowestPoint+2)
	fmt.Printf("Result: %d\n", result)
}
