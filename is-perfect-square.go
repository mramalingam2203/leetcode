package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}

	left, right := 2, num/2

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == num {
			return true
		} else if square < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	num := 16
	result := isPerfectSquare(num)
	fmt.Printf("%d is a perfect square: %v\n", num, result) // Output: "16 is a perfect square: true"
}
