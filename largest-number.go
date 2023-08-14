// https://leetcode.com/problems/largest-number/

package main

import (
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) < 1 || len(nums) > 100 {
		return "0"
	}

	s1 := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 1e9 {
			return "0"
		}

		s1[i] = strconv.Itoa(nums[i])
	}

}

func main() {

	array := []int{3, 30, 34, 5, 9}
	largestNumber(array)

}
