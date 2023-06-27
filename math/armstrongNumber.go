// https://leetcode.com/problems/armstrong-number

package main

import (
	"fmt"
	_ "os"
)

func power(x int, y uint32) int {

	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	}

	return x * power(x, y/2) * power(x, y/2)
}

// Function to check whether the given number is
// Armstrong number or not
func isArmstrong(x int) bool {
	// Calling order function
	n := order(x)
	temp := x
	sum := 0
	var r int
	for temp != 0 {
		r = temp % 10
		sum += power(r, uint32(n))
		temp /= 10
	}

	// If satisfies Armstrong condition
	return (sum == x)
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
	fmt.Println(isArmstrong(153))
	fmt.Println(isArmstrong(1253))

}
