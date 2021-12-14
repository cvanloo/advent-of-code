package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

var (
	cave    map[pnt]int
	flashes int
	flashed = map[pnt]int{}
	steps   = 100

	inpLines   = 10
	inpColumns = 10
)

func panicif(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(name string) (lines []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}

	lines = []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	return
}

func parseCave(lines []string) error {
	cave = make(map[pnt]int)

	for i, v := range lines {
		parts := strings.Split(v, "")
		for j, v := range parts {
			asi, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			cave[pnt{i, j}] = asi
		}
	}
	return nil
}

func printCave() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			val := pat(x, y)
			if 0 == val {
				fmt.Print(string(colorRed), val)
				fmt.Print(string(colorReset))
			} else {
				fmt.Print(val)
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}

type pnt struct {
	x int
	y int
}

func pat(x, y int) int {
	return cave[pnt{x, y}]
}

func bounds(p pnt) (ok bool) {
	_, ok = cave[p]
	return
}

func neighbors(p pnt) (pnts []pnt) {
	points := []pnt{
		{p.x - 1, p.y - 1}, // al
		{p.x - 1, p.y},     // a
		{p.x - 1, p.y + 1}, // ar
		{p.x, p.y - 1},     // l
		{p.x + 1, p.y},     // b
		{p.x + 1, p.y - 1}, // bl
		{p.x + 1, p.y + 1}, // br
		{p.x, p.y + 1},     // r
	}

	pnts = make([]pnt, 0, 8)

	for _, p := range points {
		if bounds(p) {
			pnts = append(pnts, p)
		}
	}

	return
}

func step() {
	flashed = map[pnt]int{}

	// increase all octopuses by one
	for p := range cave {
		cave[p]++
	}

	// light it up!
	for p := range cave {
		flash(p)
	}

	// reset flashed octi to zero
	for p := range flashed {
		cave[p] = 0
	}
}

func flash(p pnt) {
	val := cave[p]
	_, ok := flashed[p]

	if val > 9 && !ok {
		flashes++
		flashed[p] = val

		// increase all neighbors by one
		neigh := neighbors(p)

		for _, octo := range neigh {
			cave[octo]++
			flash(octo)
		}
	}
}

func main() {
	lines, err := readFile("input.txt")
	panicif(err)

	err = parseCave(lines)
	panicif(err)

	fmt.Println("Before any steps:")
	printCave()

	for i := 0; i < steps; i++ {
		step()
		fmt.Printf("After step %d:\n", i+1)
		printCave()
	}

	fmt.Println("Flashes:", flashes)
}
