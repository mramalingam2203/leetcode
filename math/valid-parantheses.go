// https://leetcode.com/problems/valid-parentheses/

package main

import (
	"fmt"
	"os"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func isValid(s string) bool {
	if len(s) < 1 || len(s) > 1e4 {
		fmt.Println("Invalid length of string")
		os.Exit(0)
	}

	srune := []rune(s)
	// bracket (lr) -- 40, 41
	// bracket (lr) -- 91, 93
	// brace (lr) -- 123, 125
	fmt.Println(srune)

	/*
		for i := 0; i < len(srune); i++ {
			if srune[i] != 40 || srune[i] != 41 || srune[i] != 91 || srune[i] != 93 || srune[i] != 123 || srune[i] != 125 {
				fmt.Println(srune[i])
				fmt.Println("Invalid character in string")
				os.Exit(0)
			}
		}
	*/

	var stack Stack

}

func main() {
	fmt.Println(isValid("(){}[]"))

}
