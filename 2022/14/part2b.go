package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func parseInput(input string) (fields map[Point]Structure, lowestPoint, mostLeft, mostRight int, err error) {
	lowestPoint, mostLeft, mostRight = 0, math.MaxInt, 0
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

			if startX == endX {
				if startY > endY {
					startY, endY = endY, startY
				}
				for i := startY; i <= endY; i++ {
					fields[Point{startX, i}] = Wall
				}
			} else if startY == endY {
				if startX > endX {
					startX, endX = endX, startX
				}
				for i := startX; i <= endX; i++ {
					fields[Point{i, startY}] = Wall
				}
			} else {
				panic("diagonals are not allowed")
			}

			if endY > lowestPoint {
				lowestPoint = endY
			}
			if endX > mostRight {
				mostRight = startX
			}
			if startX < mostLeft {
				mostLeft = startX
			}
		}
	}
	return
}

func simulateSand(fields map[Point]Structure, spawnPoint Point, lowestPoint, mostLeft, mostRight int) uint {
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

			fields[Point{x, y}] = Sand
			printField(fields, lowestPoint, mostLeft, mostRight)
			fields[Point{x, y}] = Air
		}

		fields[Point{x, y}] = Sand
		sandAtRest += 1
		printField(fields, lowestPoint, mostLeft, mostRight)

		if spawnPoint.Equal(Point{x, y}) {
			break
		}
	}

	return sandAtRest
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func printField(fields map[Point]Structure, lowestPoint, mostLeft, mostRight int) {
	fmt.Print("\033[H\033[2J")
	for y := 0; y <= lowestPoint; y++ {
		for x := mostLeft - 10; x <= mostRight+10; x++ {
			if x == 500 && y == 0 {
				fmt.Print(Red + "+" + Reset)
				continue
			}
			if y == lowestPoint {
				fmt.Print(Purple + "█" + Reset)
				continue
			}
			switch (fields[Point{x, y}]) {
			case Air:
				fmt.Print(Blue + "." + Reset)
			case Wall:
				fmt.Print(Purple + "#" + Reset)
			case Sand:
				fmt.Print(Yellow + "o" + Reset)
			default:
				fmt.Print(Red + "?" + Reset)
			}
		}
		fmt.Println()
	}
	time.Sleep(time.Millisecond * 50)
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	panicIf(err)
	fields, lowestPoint, mostLeft, mostRight, err := parseInput(string(input))
	panicIf(err)
	fmt.Printf("Number of wall fields parsed: %d\n", len(fields))
	fmt.Printf("Lowest: %d, Left: %d, Right: %d\n", lowestPoint+2, mostLeft, mostRight)
	result := simulateSand(fields, Point{500, 0}, lowestPoint+2, mostLeft, mostRight)
	fmt.Printf("Result: %d\n", result)
}
