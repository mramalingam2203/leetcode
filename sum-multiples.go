// https://leetcode.com/problems/sum-multiples/

package main

import (
	"fmt"
	"strconv"
)

func sumOfMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if divisibility(i) == true {
			sum += i
		}
	}

	return sum
}

func main() {
	N := 9
	fmt.Println(sumOfMultiples(N))

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

	return divisibilityBy7(n)
}

func divisibilityBy7(n int) bool {

	if n == 0 || n == 7 {
		return true
	}

	if n < 7 {
		return false
	}

	return divisibilityBy7(n - 7)

}
