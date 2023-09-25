// https://leetcode.com/problems/add-digits/

package main

import "fmt"

func main() {
	fmt.Println(addDigits(38))
}

func addDigits(num int) int {
	for num >= 10 {
		digit_sum := 0

		for num > 0 {
			digit := num % 10
			digit_sum += digit
			num /= 10
		}

		num = digit_sum
	}
	return num
}
