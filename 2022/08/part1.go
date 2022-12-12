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

type VisibilityChecker struct {
	visible bool
}

func (v *VisibilityChecker) Check(tree int, trees []int) {
	if !v.visible {
		v.visible = checkVisibility(tree, trees)
	}
}

func checkVisibility(tree int, trees []int) bool {
	for _, t := range trees {
		if t >= tree {
			return false
		}
	}
	return true
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

	visibilityCount := 2*rows + 2*(cols-2)

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

			visibility := &VisibilityChecker{false}

			visibility.Check(current, toLeft)
			visibility.Check(current, toRight)
			visibility.Check(current, toTop)
			visibility.Check(current, toBottom)

			if visibility.visible {
				visibilityCount++
			}
		}
	}

	fmt.Printf("%d trees are visible\n", visibilityCount)
}
