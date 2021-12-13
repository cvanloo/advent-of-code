package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func parseLine(line string) (ps []paren) {
	ps = make([]paren, 0, len(line))
	for _, r := range line {
		ps = append(ps, parseParen(r))
	}
	return
}

type paren int

const (
	oParen paren = iota
	cParen
	oBrace
	cBrace
	oBrack
	cBrack
	oLt
	cGt
)

func parseParen(token rune) paren {
	switch token {
	case '(':
		return oParen
	case ')':
		return cParen
	case '{':
		return oBrace
	case '}':
		return cBrace
	case '[':
		return oBrack
	case ']':
		return cBrack
	case '<':
		return oLt
	case '>':
		return cGt
	default:
		panic("unreachable")
	}
}

type stack struct {
	contents []paren
}

func (s *stack) push(p paren) {
	s.contents = append(s.contents, p)
}

func (s *stack) pop() (p paren) {
	l := len(s.contents)
	if l > 0 {
		p = s.contents[l-1]
		s.contents = s.contents[:l-1]
	}
	return
}

func (s *stack) next() bool {
	l := len(s.contents)
	return l > 0
}

func compile(tokens []paren) (ps *stack, ok bool) {
	ps = &stack{make([]paren, 0)}

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if tok%2 == 0 {
			// opening token
			ps.push(tok)
		} else {
			// closing token
			prev := ps.pop()
			if prev != tok-1 {
				return ps, false
			}
		}
	}
	return ps, true
}

func complete(ps *stack) (cmp []paren) {
	cmp = make([]paren, 0, len(ps.contents))

	for ps.next() {
		tok := ps.pop()
		next := tok + 1
		cmp = append(cmp, next)
	}
	return
}

func points(p paren) int {
	switch p {
	case cParen:
		return 1
	case cBrack:
		return 2
	case cBrace:
		return 3
	case cGt:
		return 4
	default:
		panic("unreachable")
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	results := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total := 0
		line := scanner.Text()
		tokens := parseLine(line)
		unf, ok := compile(tokens)
		if ok {
			fin := complete(unf)
			for _, f := range fin {
				total = total*5 + points(f)
			}
			results = append(results, total)
		}
	}

	// sort results by size
	sort.Slice(results, func(i, j int) bool { return results[i] < results[j] })

	// find middle score (there are always an odd number of scores)
	middle := len(results) / 2
	fmt.Printf("Score: %d\n", results[middle])
}
