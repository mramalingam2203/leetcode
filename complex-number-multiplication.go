// https://leetcode.com/problems/complex-number-multiplication/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	complexNumberMultiply("+1+2i", "1-2i")
}

func complexNumberMultiply(num1 string, num2 string) string {

	// n1 := make([]int, 0)
	splitNumber(num1)
	return "hi"

}

func splitNumber(s string) []int {

	signs := make([]string, 0)
	for char := range s {
		if s[char] == '+' || s[char] == '-' {
			signs = append(signs, string(s[char]))
		}
	}

	fmt.Println(signs)

	s1 := strings.Split(s, "-")
	if len(s1) < 2 {
		s1 = strings.Split(s, "+")
	}
	fmt.Println(s1)
	result := make([]int, 2)
	result[0], _ = strconv.Atoi(s1[0])
	result[1], _ = strconv.Atoi(strings.Split(s1[1], "")[0])
	return result

}
