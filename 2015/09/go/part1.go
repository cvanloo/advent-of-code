package main

import (
	"io"
	"math"
	"os"
)

type leg struct {
	target   *node
	distance int
}

type node struct {
	name       string
	distance   int
	neighbours []*leg
}

func parseInput(input string) ([]*node, error) {
	return nil, nil
}

func nodesToSet(nodes []*node) map[*node]struct{} {
	var nodesSet map[*node]struct{}
	for _, n := range nodes {
		var nCopy node = *n
		nodesSet[&nCopy] = struct{}{}
	}
	return nodesSet
}

func shortestPath(unvisited map[*node]struct{}, start *node, end *node) int {
	current := start
	for {
		for _, l := range current.neighbours {
			neighbour := l.target
			distanceToNeighbour := l.distance

			newDistance := current.distance + distanceToNeighbour
			if newDistance < neighbour.distance {
				neighbour.distance = newDistance
			}
		}

		delete(unvisited, current)

		smallest := end
		for next := range unvisited {
			if next.distance < smallest.distance {
				smallest = next
			}
		}

		if smallest.distance == math.MaxInt {
			// no path from start to end
			return smallest.distance
		}

		if smallest == end {
			// we found our path
			return smallest.distance
		}

		current = smallest
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	nodes, err := parseInput(string(input))
	if err != nil {
		panic(err)
	}

	for _, startNode := range nodes {
		set := nodesToSet(nodes)
		shortestPath(set, startNode, )
	}
}
