// https://leetcode.com/problems/multiply-strings/

package main

import (
	"fmt"
	_ "os"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if num1 == "1" && num2 == "1" {
		return "1"
	}

	return strconv.Itoa(n1 * n2)

}

func mag(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a int, b int) int {
	if a <= b {
		return b
	}
	return a
}

func main() {
	n1 := "123"
	n2 := "456"
	fmt.Println(multiply(n1, n2))
}
