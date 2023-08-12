// https://leetcode.com/problems/sum-multiples/

package main

import (
	"fmt"
	"strconv"
)

func sumOfMultiples(n int) int {
	return 0
}

func main() {
	N := 28
	fmt.Println(divisibility(N))
}

func divisibility(n int) bool {
	// by 5
	str := strconv.Itoa(n)
	if str[len(str)-1] == '0' || str[len(str)-1] == '5' {
		return true
	}

	// by 3
	sum := 0
	for i := 0; i < len(str); i++ {
		k := string(str[i])
		digit, _ := strconv.Atoi(k)
		sum += digit
	}

	if sum%3 == 0 {
		return true
	}

	// by 7

	return false
}
