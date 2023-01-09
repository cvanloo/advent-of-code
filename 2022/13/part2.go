package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type Node interface {
	// Compare compares two nodes.
	// Returns a negative number if nodes are in the right order,
	// a positive number if nodes are in the wrong order, and
	// zero if order cannot be determined, in which case the next nodes need be compared.
	Compare(other Node) int
	// IsAtom returns true only if the node is of type atom.
	// Same as !IsCons()
	IsAtom() bool
	// IsCons returns true only if the node is of type cons.
	// Same as !IsAtom()
	IsCons() bool
	// Value returns the value of an atom, but will panic when called on a cons.
	Value() int
	// Children returns the children of a cons, but will panic when called on an atom.
	Children() []Node
}

type atom struct {
	value int
}

// atom implements Node
var _ Node = (*atom)(nil)

func (a *atom) Compare(other Node) int {
	if other.IsAtom() {
		// compare atom with atom
		// lower integer should come first
		return a.value - other.Value()
	} else {
		// compare atom with cons
		// convert atom to cons, retry comparison
		consAtom := &cons{
			children: []Node{a},
		}
		return consAtom.Compare(other)
	}
}

func (*atom) IsAtom() bool {
	return true
}

func (*atom) IsCons() bool {
	return false
}

func (a *atom) Value() int {
	return a.value
}

func (*atom) Children() []Node {
	panic("must not call Children() on atom")
}

func (a *atom) String() string {
	return fmt.Sprintf("%d", a.value)
}

type cons struct {
	children []Node
}

// cons implements Node
var _ Node = (*cons)(nil)

func (c *cons) Compare(other Node) int {
	if other.IsAtom() {
		// compare cons with atom
		// convert atom to cons, retry comparison
		return c.Compare(&cons{
			children: []Node{other},
		})
	} else {
		// compare cons with cons
		// compare atoms of both cons
		// if left cons runs out of items first, nodes are in right order
		idx := 0
		maxLeft := len(c.children) - 1
		maxRight := len(other.Children()) - 1
		for {
			if idx > maxLeft {
				if maxLeft == maxRight {
					return 0
				}
				return -1
			}
			if idx > maxRight {
				return +1
			}

			leftChild := c.children[idx]
			rightChild := other.Children()[idx]
			idx++

			res := leftChild.Compare(rightChild)
			if res != 0 {
				return res
			}
		}
	}
}

func (*cons) IsAtom() bool {
	return false
}

func (*cons) IsCons() bool {
	return true
}

func (*cons) Value() int {
	panic("must not call Value() on cons")
}

func (c *cons) Children() []Node {
	return c.children
}

func (c *cons) String() string {
	output := "["
	last := len(c.children) - 1
	for cur, child := range c.children {
		if cur < last {
			output += fmt.Sprintf("%s,", child)
		} else {
			output += fmt.Sprintf("%s", child)
		}
	}
	output += "]"
	return output
}

func idioticSplit(s string, delims []rune) []string {
	var out []string
	lastPos := 0
	for i, r := range s {
		for _, d := range delims {
			if r == d {
				if i > 0 {
					out = append(out, s[lastPos:i])
				}
				lastPos = i + 1
				out = append(out, string(d))
				break
			}
		}
	}
	return out
}

func parseInput(input string) ([]Node, error) {
	input = strings.TrimRight(input, "\n\r")
	lines := strings.Split(input, "\n")

	topNodes := make([]Node, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimRight(line, "\r\n")

		newCons := &cons{
			children: []Node{},
		}

		previousCons := []*cons{nil, newCons}
		lastCons, currentCons := (*cons)(nil), newCons

		fields := idioticSplit(line, []rune{',', '[', ']'})

		for _, field := range fields {
			switch field {
			case "[":
				// Start new cons
				lastCons = currentCons
				currentCons = &cons{}
				previousCons = append(previousCons, currentCons)

				lastCons.children = append(lastCons.children, currentCons)
			case "]":
				// End current cons
				previousCons = previousCons[:len(previousCons)-1] // remove current cons
				currentCons = lastCons                            // last is new current
				lastCons = previousCons[len(previousCons)-2]      // second-to-last is new last
			case ",":
				fallthrough
			case "":
				continue
			default:
				// Parse int, add to current cons
				val, err := strconv.Atoi(field)
				panicIf(err)
				currentCons.children = append(currentCons.children, &atom{val})
			}
		}

		if len(newCons.children) == 0 {
			// empty input line
			continue
		}
		if len(previousCons) > 2 {
			return topNodes, errors.New("unfinished cons in line")
		}
		if len(newCons.children) > 1 {
			return topNodes, errors.New("too many children in one line")
		}
		topNodes = append(topNodes, newCons.children[0])
	}

	return topNodes, nil
}

func getResultText(result int) string {
	var output string
	if result < 0 {
		output = "correct order"
	} else if result > 0 {
		output = "incorrect order"
	} else {
		output = "no conclusion possible"
	}

	return fmt.Sprintf("Result (%d): %s", result, output)
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	panicIf(err)

	lists, err := parseInput(string(input))
	if len(lists)%2 != 0 {
		log.Fatal("want even number of lists")
	}
	panicIf(err)

	divider1 := &cons{[]Node{&cons{[]Node{&atom{2}}}}}
	divider2 := &cons{[]Node{&cons{[]Node{&atom{6}}}}}
	lists = append(lists, divider1)
	lists = append(lists, divider2)

	sort.Slice(lists, func(i, j int) bool {
		left := lists[i]
		right := lists[j]
		res := left.Compare(right)
		return res < 0
	})

	var d1, d2 int
	for idx, node := range lists {
		idx1 := idx + 1
		if node == divider1 {
			d1 = idx1
		} else if node == divider2 {
			d2 = idx1
		}
	}
	fmt.Printf("The final result is %d\n", d1*d2)
}
