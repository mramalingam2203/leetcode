package main

import (
	"fmt"
)

type stack []float64

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, float64) {
	// FIXME: What do we do if the stack is empty, though?

	l := len(s)
	return s[:l-1], s[l-1]
}

func main() {
	s := make(stack, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	s, p := s.Pop()
	fmt.Println(p)

}
