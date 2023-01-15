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
			nodes = append(nodes, endNode)
		}

		if startNode == nil {
			startNode = &node{
				name:     startCityName,
				distance: math.MaxInt,
			}
			nodes = append(nodes, startNode)
		}

		startNode.neighbours = append(startNode.neighbours, leg{
			target:   endNode,
			distance: distanceBetween,
		})

		endNode.neighbours = append(endNode.neighbours, leg{
			target:   startNode,
			distance: distanceBetween,
		})
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

func shortestPath(unvisited map[*node]struct{}, start *node, end *node) (path []node, distance int) {
	path = make([]node, 0, len(unvisited))
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
		path = append(path, *current)

		if current == end {
			// we found our path
			return path, current.distance
		}

		smallest := end
		for next := range unvisited {
			if next.distance < smallest.distance {
				smallest = next
			}
		}

		if smallest.distance == math.MaxInt {
			// no path from start to end
			return path, smallest.distance
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

	for _, n := range nodes {
		fmt.Printf("%s -> ", n.name)
		for _, neighbour := range n.neighbours {
			fmt.Printf("%s ", neighbour.target.name)
		}
		fmt.Println()
	}

	for _, startNode := range nodes {
		for _, endNode := range nodes {
			if startNode == endNode {
				continue
			}
			set := nodesToSet(nodes)
			startNode.distance = 0
			path, distance := shortestPath(set, startNode, endNode)
			//if distance != math.MaxInt && len(path) == len(nodes) {
			fmt.Printf("start %s, end %s: ", startNode.name, endNode.name)
			for _, p := range path {
				fmt.Printf("%s, ", p.name)
			}
			fmt.Printf("-> distance: %d\n", distance)
			//}
		}
	}
}
