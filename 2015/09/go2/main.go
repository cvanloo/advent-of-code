package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func copy[T any](a, b []T) {
	for i := range a {
		b[i] = a[i]
	}
}

func delete[T any](a []T, i int) []T {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func perm(n []int) (ps [][]int) {
	if len(n) == 1 {
		return [][]int{n}
	}
	for i := range n {
		selected := n[i]
		left := make([]int, len(n))
		copy(n, left)
		left = delete(left, i)
		nextps := perm(left)
		for _, nextp := range nextps {
			p := append([]int{selected}, nextp...)
			ps = append(ps, p)
		}
	}
	return ps
}

type distance struct {
	start, end string
	cost int
}

func findDistance(distances []distance, a, b string) int {
	for _, p := range distances {
		if (p.start == a && p.end == b) || (p.start == b && p.end == a) {
			return p.cost
		}
	}
	panic("not found")
}

func appendUnique[T string|int](a []T, els ...T) []T {
outer:
	for _, el := range els {
		for i := range a {
			if a[i] == el {
				continue outer
			}
		}
		a = append(a, el)
	}
	return a
}

func main2() {
	fmt.Printf("%+v\n", perm([]int{}))
	fmt.Printf("%+v\n", perm([]int{0}))
	fmt.Printf("%+v\n", perm([]int{0, 1, 2}))
	fmt.Printf("%+v\n", perm([]int{0, 1, 2, 3}))
}

func main() {
	//bs := must(os.ReadFile("sample.txt"))
	bs := must(os.ReadFile("input.txt"))
	content := string(bs)
	lines := strings.Split(content, "\n")
	distances := make([]distance, len(lines))
	places := []string{}
	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) < 5 {
			continue
		}
		d := distance{
			start: parts[0],
			end: parts[2],
			cost: must(strconv.Atoi(parts[4])),
		}
		distances[i] = d
		places = appendUnique(places, d.start, d.end)
	}
	fmt.Printf("%+v\n", distances)
	fmt.Printf("%+v\n", places)
	indecies := make([]int, len(places))
	for i := range indecies {
		indecies[i] = i
	}
	permutations := perm(indecies)
	fmt.Printf("%+v\n", permutations)

	shortestCost, shortestVector := math.MaxInt, []int{}
	longestCost, longestVector := 0, []int{}
	for _, perm := range permutations {
		totalCost := 0
		for i := 1; i < len(perm); i++ {
			a := perm[i-1]
			b := perm[i]
			totalCost += findDistance(distances, places[a], places[b])
		}
		if totalCost < shortestCost {
			shortestCost = totalCost
			shortestVector = perm
		}
		if totalCost > longestCost {
			longestCost = totalCost
			longestVector = perm
		}
	}
	fmt.Printf("Part 1: %+v (%+v)\n", shortestCost, shortestVector)
	fmt.Printf("Part 2: %+v (%+v)\n", longestCost, longestVector)
}
