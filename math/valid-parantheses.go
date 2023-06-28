// https://leetcode.com/problems/valid-parentheses/

package main

import (
	"fmt"
	"os"
)

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

	stack := []rune{}
	bracketsMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != bracketsMap[char] {
				return false
			}
			stack = stack[:len(stack)-1] // Pop the last element
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("(){}[}"))

}
