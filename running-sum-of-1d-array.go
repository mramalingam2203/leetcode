// https://leetcode.com/problems/running-sum-of-1d-array/

package main

import (
	"fmt"
	"os"
)

func main() {
	list := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(list))
}

func runningSum(nums []int) []int {
	if len(nums) < 1 || len(nums) > 1000 {
		fmt.Println("invalid inputs")
		os.Exit(0)
	}
	sums := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] < -1e6 || nums[i] > 1e6 {
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
