package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func scoreVisibility(tree int, trees []int) (score int) {
	for _, t := range trees {
		score++
		if t >= tree {
			break
		}
	}
	return score
}

func reverse(ints []int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func main() {
	content, err := os.ReadFile(os.Args[1])
	panicIf(err)

	input := string(content)
	lines := strings.Fields(input)

	firstLine := lines[0]
	rows := len(lines)
	cols := len(firstLine)

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	for col, line := range lines {
		for row, r := range line {
			val, _ := strconv.Atoi(string(r))
			grid[col][row] = val
		}
	}

	maxScore := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			current := grid[r][c]
			toLeft := grid[r][:c]
			toRight := grid[r][c+1:]

			var toTop []int
			for top := r - 1; top >= 0; top-- {
				t := grid[top][c]
				toTop = append(toTop, t)
			}

			var toBottom []int
			for bot := r + 1; bot < rows; bot++ {
				b := grid[bot][c]
				toBottom = append(toBottom, b)
			}

			toLeftReversed := make([]int, len(toLeft))
			copy(toLeftReversed, toLeft)
			reverse(toLeftReversed)

			s1 := scoreVisibility(current, toLeftReversed)
			s2 := scoreVisibility(current, toRight)
			s3 := scoreVisibility(current, toTop)
			s4 := scoreVisibility(current, toBottom)

			score := s1 * s2 * s3 * s4
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Printf("The highest score is %d\n", maxScore)
}
