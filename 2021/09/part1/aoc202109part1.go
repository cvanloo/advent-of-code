package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	area = [inputLength][inputWidth]int{}
	risk = 0
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
}

func getAbove(x, y int) int {
	x--
	if x >= 0 {
		val := area[x][y]
		return val
	}

	// In a more serious program you probably would do something like `return
	// 0, false`
	return math.MaxInt
}

func getBelow(x, y int) int {
	x++
	if x < inputLength {
		val := area[x][y]
		return val
	}
	return math.MaxInt
}

func getLeft(x, y int) int {
	y--
	if y >= 0 {
		val := area[x][y]
		return val
	}
	return math.MaxInt
}

func getRight(x, y int) int {
	y++
	if y < inputWidth {
		val := area[x][y]
		return val
	}
	return math.MaxInt
}

func calculateRisk(height int) {
	risk += height + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	parseInput(file)
	file.Close()

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
		}
	}

	fmt.Println("The risk is:", risk)
}
