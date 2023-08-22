// https://leetcode.com/problems/

package main

import "unicode"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item
}

func (s *Stack) Top() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func evalRPN(tokens []string) int {
	stack := Stack{}

	for token := 0; token < len(tokens); token++ {
		runeArray := []rune(tokens[token])
		if unicode.IsDigit(runeArray[0]) {
			stack.Push(runeArray[0])
		}
	}

	return 0
}

func main() {
	toks := []string{"2", "1", "+", "3", "*"}
	evalRPN(toks)

}
