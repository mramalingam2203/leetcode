// https://leetcode.com/problems/running-sum-of-1d-array/

package main

import (
	"fmt"
	"os"
)

func main() {
	list := []int{10, 4, 8, 3}
	leftRightDifference(list)
}

func leftRightDifference(nums []int) []int {
	if len(nums) < 1 || len(nums) > 1000 {
		fmt.Println("invalid inputs")
		os.Exit(0)
	}

	sums := cumulativeSum(nums)
	fmt.Println(sums)
	n := len(sums)

	for i := 1; i < n-1; i++ {
		left := sums[0] - sums[i-1]
		right := sums[i] - sums[n-1]
		fmt.Print(left-right, " ")

	}

	return nums

}

func cumulativeSum(nums []int) []int {
	sums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > 1e5 {
			fmt.Println("invalid inputs")
			os.Exit(0)
		}

		if i != 0 {
			sums[i] = sums[i-1] + nums[i]
		} else {
			sums[i] = nums[i]
		}
	}

	return sums
}
