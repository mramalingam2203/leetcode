// https://leetcode.com/problems/

package main

import "unicode"

type Stack struct {
	items []byte
}

func (s *Stack) Push(char byte) {
	s.items = append(s.items, char)
}

func (s *Stack) IsEmpty() bool {

	return len(s.items) == 0

}

func (s *Stack) Pop() byte {
	if len(s.items) == 0 {
		return '0'
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item

}

func evalRPN(tokens []string) int {
	//stack := Stack{}

	for token := 0; token < len(tokens); token++ {
		if unicode.IsDigit(tokens[token]) {
			stack.Push(tokens[token])
		}
	}

}

func main() {
	toks := []string{"2", "1", "+", "3", "*"}
	evalRPN(toks)

}
