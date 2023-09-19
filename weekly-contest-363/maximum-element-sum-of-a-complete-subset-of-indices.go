// https://leetcode.com/contest/weekly-contest-363/problems/maximum-element-sum-of-a-complete-subset-of-indices/

package main

import "fmt"

func main() {
	array := []int{8, 7, 3, 5, 7, 2, 4, 9}
	maximumSum(array)
	fmt.Println(isPerfectSquare(36))

}

func maximumSum(nums []int) int64 {

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPerfectSquare(nums[i]*nums[j]) == true {
				fmt.Println(nums[i], nums[j])
			}
		}
	}

	return 10
}

func isPerfectSquare(n int) bool {

	// Base case: 0 and 1 are perfect squares
	if n <= 1 {
		return true
	}

	var left, right uint64 = 1, uint64(n)
	var mid, square uint64

	for left <= right {
		mid = left + (right-left)/2

		square = mid * mid

		if square == uint64(n) {
			return true
		} else if square < uint64(n) {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}
