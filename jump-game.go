/*
Merge Intervals
Insert Interval
Spiral Matrix II
Unique Paths II
Minimum Path Sum
Set Matrix Zeroes
Search a 2D Matrix
Sort Colors
*/

package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) < 1 || len(nums) > 1e4 {
		return false
	}

	furthestIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 1e5 {
			return false
		}
		if i > furthestIndex {
			return false
		}
		furthestIndex = max(furthestIndex, i+nums[i])
	}

	return furthestIndex >= len(nums)-1
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	a := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(a))

}
