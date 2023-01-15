package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type leg struct {
	target   *node
	distance int
}

type node struct {
	name       string
	distance   int
	neighbours []leg
}

func findCityByName(nodes []*node, name string) *node {
	for _, n := range nodes {
		if n.name == name {
			return n
		}
	}
	return nil
}

func parseInput(input string) ([]*node, error) {
	var nodes []*node

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		words := strings.Fields(line)
		startCityName := words[0]
		endCityName := words[2]
		distanceBetween, err := strconv.Atoi(words[4])
		if err != nil {
			panic(err)
		}

		endNode := findCityByName(nodes, endCityName)
		startNode := findCityByName(nodes, startCityName)

		if endNode == nil {
			endNode = &node{
				name:     endCityName,
				distance: math.MaxInt,
			}
		}

		if startNode == nil {
			startNode = &node{
				name:     startCityName,
				distance: math.MaxInt,
			}
		}

		startNode.neighbours = append(startNode.neighbours, leg{
			target:   endNode,
			distance: distanceBetween,
		})

		nodes = append(nodes, endNode)
		nodes = append(nodes, startNode)
	}

	return nodes, nil
}

func nodesToSet(nodes []*node) map[*node]struct{} {
	nodesSet := make(map[*node]struct{}, len(nodes))
	for _, n := range nodes {
		n.distance = math.MaxInt
		nodesSet[n] = struct{}{}
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
		for _, endNode := range nodes {
			if startNode == endNode {
				continue
			}
			set := nodesToSet(nodes)
			startNode.distance = 0
			distance := shortestPath(set, startNode, endNode)
			if distance == math.MaxInt {
				continue
			}
			fmt.Printf("Distance: %d\n", distance)
		}
	}
}
