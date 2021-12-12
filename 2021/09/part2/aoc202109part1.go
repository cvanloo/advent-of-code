package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	//inputLength = 5
	//inputWidth  = 10
	inputLength = 100
	inputWidth  = 100
)

var (
	area    = [inputLength][inputWidth]int{}
	risk    = 0
	basins  = []int{}
	visited = [inputLength][inputWidth]bool{}
)

func parseInput(file *os.File) {
	scanner := bufio.NewScanner(file)

	for i := 0; i < inputLength; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, "")

		for j := 0; j < inputWidth; j++ {
			p := parts[j]
			val, err := strconv.Atoi(p)
			if err != nil {
				log.Fatalf("invalid input: %v", err)
			}

			area[i][j] = val
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func getAbove(x, y int) int {
	x--
	if x >= 0 {
		return area[x][y]
	}
	return math.MaxInt
}

func getBelow(x, y int) int {
	x++
	if x < inputLength {
		return area[x][y]
	}
	return math.MaxInt
}

func getLeft(x, y int) int {
	y--
	if y >= 0 {
		return area[x][y]
	}
	return math.MaxInt
}

func getRight(x, y int) int {
	y++
	if y < inputWidth {
		return area[x][y]
	}
	return math.MaxInt
}

func calculateRisk(height int) {
	risk += height + 1
}

func calculateBasin(x, y int) (size int) {
	size = 1
	visited[x][y] = true

	if getBelow(x, y) < 9 && !visited[x+1][y] {
		size += calculateBasin(x+1, y)
	}
	if getAbove(x, y) < 9 && !visited[x-1][y] {
		size += calculateBasin(x-1, y)
	}
	if getRight(x, y) < 9 && !visited[x][y+1] {
		size += calculateBasin(x, y+1)
	}
	if getLeft(x, y) < 9 && !visited[x][y-1] {
		size += calculateBasin(x, y-1)
	}

	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	parseInput(file)
	file.Close()

	// find lowpoints
	for i := 0; i < inputLength; i++ {
		for j := 0; j < inputWidth; j++ {
			cur := area[i][j]
			if cur >= getLeft(i, j) {
				continue
			}
			if cur >= getAbove(i, j) {
				continue
			}
			if cur >= getRight(i, j) {
				continue
			}
			if cur >= getBelow(i, j) {
				continue
			}
			calculateRisk(cur)
			basins = append(basins, calculateBasin(i, j))
		}
	}

	// sort to get largest first
	sort.Slice(basins, func(i, j int) bool { return basins[i] > basins[j] })

	fmt.Println("The risk is:", risk)
	fmt.Println("Basins:", basins[0]*basins[1]*basins[2])
}
