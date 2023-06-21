package main

import "fmt"

func generateParentheses(n int) []string {
	result := []string{}
	generateParenthesesHelper(n, n, "", &result)
	return result
}

func generateParenthesesHelper(leftCount, rightCount int, currentStr string, result *[]string) {
	if leftCount == 0 && rightCount == 0 {
		*result = append(*result, currentStr)
		return
	}

	if leftCount > 0 {
		generateParenthesesHelper(leftCount-1, rightCount, currentStr+"(", result)
	}

	if rightCount > leftCount {
		generateParenthesesHelper(leftCount, rightCount-1, currentStr+")", result)
	}
}

func main() {
	n := 3
	combinations := generateParentheses(n)
	fmt.Println(combinations)
}
