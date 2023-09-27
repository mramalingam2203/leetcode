// https://leetcode.com/problems/remove-invalid-parentheses/

package main

func main() {

	removeInvalidParentheses("[3]He[llo]")
}

func removeInvalidParentheses(s string) []string {

}

func isValid(s string) bool {

	count := 0

	for char := range s {

		if char == '(' {
			count++
		} else if char == ')' {
			count--
			if count < 0 {
				return false
			}

		}
	}
	return count == 0
}
