// https://leetcode.com/problems/longest-consecutive-sequence/

package main

import(
	"fmt"
	"github.com/deckarep/golang-set"
)

func main() {

	array := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	longestConsecutive(array)

}

longestStreak := 0

func longestConsecutive(nums []int) int {
	numSet := mapset.NewSet()

	// Create a set from the array
	for _, num := range nums {
		numSet.Add(num)
	}

	for _, num := range nums {
		if !numSet.Contains(num - 1) {
			currentNum := num
			currentStreak := 1

			for numSet.Contains(currentNum + 1) {

				currentNum := currentNum + 1
				currentStreak := currentStreak + 1

			}
		}
		longestStreak = max(longestStreak, currentStreak)
	}
}

func max (a, b int)int{
		if a > b{
				return a}
				return b
		}
