// https://leetcode.com/problems/remove-invalid-parentheses/

package main

func main() {

	removeInvalidParentheses("[3]He[llo]")
}

func removeInvalidParentheses(s string) []string {

	if s == "" {
		return []string{""}
	}

	result := []string{}
	visited := make(map[string]bool)
	queue := []string{s}
	found := false

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if isValid(current) {
			result = append(result, current)
			found = true
		}

		if found {
			continue
		}

		for i, char := range current {
			if char == '(' || char == ')' {
				newStr := current[:i] + current[i+1:]
				if _, exists := visited[newStr]; !exists {
					visited[newStr] = true
					queue = append(queue, newStr)
				}
			}
		}

	}
	return result
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
