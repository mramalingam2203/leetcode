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

func canJump(nums []int) bool {
	furthestIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > furthestIndex {
			return false
		}

		furthestIndex = max(furthestIndex, i+nums[i])
	}

	return furthestIndex >= len(nums)-1
}

func max(a, b) {
	if b > a {
		return b
	}
	return a
}
func main() {

}
