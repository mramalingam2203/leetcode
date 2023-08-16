// https://leetcode.com/problems/factorial-trailing-zeroes/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(trailingZeroes(0))
}

func trailingZeroes(n int) int {

	if n < 0 || n > 1e4 {
		return 0
	}
	fact := factorial(n)
	s := strconv.Itoa(fact)
	count := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "0" {
			count++
		}

	}
	return count

}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}
