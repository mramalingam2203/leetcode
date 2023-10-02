// https://leetcode.com/problems/count-numbers-with-unique-digits/

package main

import "fmt"

func countNumbersWithUniqueDigits(n int) int {

	if n == 0 {
		return 1
	}

	if n == 1 {
		return 10
	}

	// Initialize counts for 1-digit and 2-digit numbers
	count_1_digit := 10
	count_2_digits := 9 * 9

	// Calculate counts for numbers with 3 or more digits using dynamic programming
	for i := 2; i < n; i++ {
		count_1_digit *= (11 - i)
		count_2_digits *= (10 - i)
	}

	// Total count is the sum of 1-digit and 2-digit counts
	return count_1_digit + count_2_digits

}

func main() {

	fmt.Println(countNumbersWithUniqueDigits(2))
}
