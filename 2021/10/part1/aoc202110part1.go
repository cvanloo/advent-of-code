package main

import (
	"bufio"
	"fmt"
	"os"
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

func compile(tokens []paren) (paren, bool) {
	ps := &stack{make([]paren, 0)}

	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		if tok%2 == 0 {
			// opening token
			ps.push(tok)
		} else {
			// closing token
			prev := ps.pop()
			if prev != tok-1 {
				return tok, true
			}
		}
	}
	return 0, false
}

func points(p paren) int {
	switch p {
	case cParen:
		return 3
	case cBrack:
		return 57
	case cBrace:
		return 1197
	case cGt:
		return 25137
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

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := parseLine(line)
		inv, res := compile(tokens)
		if res {
			// got invalid token
			result += points(inv)
		}
	}

	fmt.Println("The result is:", result)
}
