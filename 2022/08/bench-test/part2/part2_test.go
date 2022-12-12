package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"
)

//go:embed input.txt
var input string

func BenchmarkAllocAndInitInOne(b *testing.B) {
	benchFunc := func() {
		lines := strings.Fields(input)

		firstLine := lines[0]
		rows := len(lines)
		cols := len(firstLine)

		grid := make([][]int, rows)
		for row, line := range lines {
			grid[row] = make([]int, cols)
			for col, r := range line {
				val, _ := strconv.Atoi(string(r))
				grid[row][col] = val
			}
		}
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkAllocAndInitSeparate(b *testing.B) {
	benchFunc := func() {
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
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkTwoGrids(b *testing.B) {
	benchFunc := func() {
		lines := strings.Fields(input)

		firstLine := lines[0]
		rows := len(lines)
		cols := len(firstLine)

		grid := make([][]int, rows)
		for row := range grid {
			grid[row] = make([]int, cols)
		}

		gridH := make([][]int, cols)
		for col := range gridH {
			gridH[col] = make([]int, rows)
		}

		for row, line := range lines {
			for col, r := range line {
				val, _ := strconv.Atoi(string(r))
				grid[row][col] = val
				gridH[col][row] = val
			}
		}

		maxScore := 0

		for r := 1; r < rows-1; r++ {
			for c := 1; c < cols-1; c++ {
				current := grid[r][c]
				toLeft := grid[r][:c]
				toRight := grid[r][c+1:]

				toTop := gridH[c][:r]
				toBottom := gridH[c][r+1:]

				toLeftReversed := make([]int, len(toLeft))
				copy(toLeftReversed, toLeft)
				reverse(toLeftReversed)

				toTopReversed := make([]int, len(toTop))
				copy(toTopReversed, toTop)
				reverse(toTopReversed)

				s1 := scoreVisibility(current, toLeftReversed)
				s2 := scoreVisibility(current, toRight)
				s3 := scoreVisibility(current, toTopReversed)
				s4 := scoreVisibility(current, toBottom)

				score := s1 * s2 * s3 * s4
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}

func BenchmarkOneGrid(b *testing.B) {
	benchFunc := func() {
		lines := strings.Fields(input)

		firstLine := lines[0]
		rows := len(lines)
		cols := len(firstLine)

		grid := make([][]int, rows)
		for row := range grid {
			grid[row] = make([]int, cols)
		}

		gridH := make([][]int, cols)
		for col := 0; col < cols; col++ {
			gridH[col] = make([]int, rows)
		}

		for row, line := range lines {
			for col, r := range line {
				val, _ := strconv.Atoi(string(r))
				grid[row][col] = val
				gridH[col][row] = val
			}
		}

		maxScore := 0

		for r := 1; r < rows-1; r++ {
			for c := 1; c < cols-1; c++ {
				current := grid[r][c]
				toLeft := grid[r][:c]
				toRight := grid[r][c+1:]

				toTop := gridH[c][:r]
				toBottom := gridH[c][r+1:]

				toLeftReversed := make([]int, len(toLeft))
				copy(toLeftReversed, toLeft)
				reverse(toLeftReversed)

				toTopReversed := make([]int, len(toTop))
				copy(toTopReversed, toTop)
				reverse(toTopReversed)

				s1 := scoreVisibility(current, toLeftReversed)
				s2 := scoreVisibility(current, toRight)
				s3 := scoreVisibility(current, toTopReversed)
				s4 := scoreVisibility(current, toBottom)

				score := s1 * s2 * s3 * s4
				if score > maxScore {
					maxScore = score
				}
			}
		}
	}

	for n := 0; n < b.N; n++ {
		benchFunc()
	}
}
