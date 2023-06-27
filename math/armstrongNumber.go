// https://leetcode.com/problems/armstrong-number

package main

import (
	"fmt"
	_ "os"
)

func power(x int, y uint64) int {

	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	}

	return x * power(x, y/2) * power(x, y/2)
}

/* Function to calculate order of the number */
func order(x int) int {
	n := 0
	for x != 0 {
		n++
		x = x / 10
	}
	return n
}

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(order(10))
}
