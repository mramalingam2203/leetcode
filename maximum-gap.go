// https://leetcode.com/problems/maximum-gap/

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 6, 9, 1}
	maximumGap(arr)
}

func maximumGap(nums []int) int {

	sort.Ints(nums)
	n := len(nums)

	if n < 1 || n > 1e5 {
		return 0
	}

	maxGap := 0

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return nums[1] - nums[0]
	}

	if n == 3 {
		if nums[1]-nums[0] > nums[2]-nums[1] {
			return nums[1] - nums[0]
		} else {
			return nums[2] - nums[1]
		}
	}

	for i := 1; i < n-1; i++ {
		if nums[i] < 0 || nums[i] > 1e9 {
			return 0
		}
		if nums[i]-nums[i-1] < nums[i+1]-nums[i] {
			maxGap = nums[i+1] - nums[i]
		}
	}

	fmt.Println("maxGap: ", maxGap)

	return maxGap
}
