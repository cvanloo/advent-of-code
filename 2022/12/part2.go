package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p Point) Neighbours() (neighbours [4]Point) {
	offsets := []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for i, offset := range offsets {
		neighbour := p.Add(offset)
		neighbours[i] = neighbour
	}

	return neighbours
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Node struct {
	point      Point
	neighbours []*Node // Distance between two nodes is always 1
	distance   uint
}

func getValue(labyrinth [][]rune, p Point) (value rune, err error) {
	if p.y < 0 || p.y >= len(labyrinth) {
		return value, fmt.Errorf("y-Coordinate out of bounds: %d", p.y)
	}
	if p.x < 0 || p.x >= len(labyrinth[0]) {
		return value, fmt.Errorf("x-Coordinate out of bounds: %d", p.x)
	}

	value = labyrinth[p.y][p.x]
	return value, nil
}

func getAllowedDirections(labyrinth [][]rune, currentPosition Point, directions []Point) (allowedDirections []Point) {
	currentValue, _ := getValue(labyrinth, currentPosition)

	for _, direction := range directions {
		targetValue, err := getValue(labyrinth, direction)
		if err != nil {
			continue
		}

		difference := targetValue - currentValue
		if difference <= 1 {
			allowedDirections = append(allowedDirections, direction)
		}
	}
	return allowedDirections
}

func createGraph(labyrinth [][]rune, startPoint, endPoint Point) (startNodes []*Node, end *Node, nodes map[*Node]struct{}) {
	nodes = map[*Node]struct{}{}
	startNodes = []*Node{}
	nodePoints := map[Point]*Node{}
	x, y := 0, 0

	for {
		if x >= len(labyrinth[0]) {
			y++
			x = 0
		}
		if y >= len(labyrinth) {
			break
		}

		currentPoint := Point{x, y}
		neighbourPoints := currentPoint.Neighbours()
		allowedNeighbours := getAllowedDirections(labyrinth, currentPoint, neighbourPoints[:])

		if len(allowedNeighbours) <= 0 {
			break
		}

		var neighbours []*Node
		for _, point := range allowedNeighbours {
			neighbourNode, ok := nodePoints[point]
			if !ok {
				neighbourNode = &Node{
					point:      point,
					neighbours: nil,
					distance:   math.MaxInt64,
				}
				nodePoints[point] = neighbourNode
			}

			neighbours = append(neighbours, neighbourNode)
		}

		currentNode, ok := nodePoints[currentPoint]
		if !ok {
			currentNode = &Node{
				point:    currentPoint,
				distance: math.MaxUint64,
			}
			nodePoints[currentPoint] = currentNode
		}
		currentNode.neighbours = neighbours

		nodes[currentNode] = struct{}{}
		if value, err := getValue(labyrinth, currentPoint); err == nil && value == 'a' {
			startNodes = append(startNodes, currentNode)
		}

		x++
	}

	//nodePoints[startPoint].distance = 0
	return startNodes, nodePoints[endPoint], nodes
}

func parseFields(input string) (fields [][]rune, start, end Point) {
	input = strings.TrimRight(input, "\r\n") // Because Windows sucks.
	lines := strings.Split(input, "\n")
	ymax, xmax := len(lines), len(lines[0])

	var startPoint, endPoint Point

	fields = make([][]rune, ymax)

	for y, line := range lines {
		line = strings.TrimRight(line, "\r") // Because Windows still sucks.
		fields[y] = make([]rune, xmax)
		for x, r := range line {
			switch r {
			case 'S':
				startPoint = Point{x, y}
				fields[y][x] = 'a'
			case 'E':
				endPoint = Point{x, y}
				fields[y][x] = 'z'
			default:
				fields[y][x] = r
			}
		}
	}

	return fields, startPoint, endPoint
}

func findShortestPath(unvisited map[*Node]struct{}, start, end *Node) (shortestDistance uint) {
	currentNode := start

	for {
		for _, neighbour := range currentNode.neighbours {
			if _, ok := unvisited[neighbour]; ok {
				distance := currentNode.distance + 1 // Edge length is always 1
				if neighbour.distance > distance {
					neighbour.distance = distance
				}
			}
		}

		delete(unvisited, currentNode)

		if currentNode == end {
			return end.distance
		}

		smallest := end
		for node := range unvisited {
			if node.distance < smallest.distance {
				smallest = node
			}
		}

		if smallest.distance == math.MaxUint64 {
			return end.distance
		}

		currentNode = smallest
	}
}

func copyMap[K, V comparable](dst, src map[K]V) {
	for k, v := range src {
		dst[k] = v
	}
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fields, startPoint, endPoint := parseFields(string(input))
	fmt.Printf("Start: %v, End: %v\n", startPoint, endPoint)

	startNodes, endNode, nodes := createGraph(fields, startPoint, endPoint)
	fmt.Printf("Startnodes: %d\n", len(startNodes))

	lengths := make([]uint, len(startNodes))
	for i, startNode := range startNodes {
		startNode.distance = 0
		unvisited := map[*Node]struct{}{}
		copyMap(unvisited, nodes)
		length := findShortestPath(unvisited, startNode, endNode)
		startNode.distance = math.MaxUint64
		lengths[i] = length
	}
	fmt.Printf("Paths: %v\n", lengths)

	sort.Slice(lengths, func(i, j int) bool {
		return lengths[i] < lengths[j]
	})

	fmt.Printf("Shortest Path: %d\n", lengths[0])
}
