// https://leetcode.com/problems/single-number-ii/

package main

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {

	if len(nums) < 1 || len(nums) > 3e4 {
		return 0
	}

	sort.Ints(nums)

	fmt.Println(nums)

	for i := 0; i < len(nums)-2; i += 3 {
		if nums[i] < -3e4 || nums[i] > 3e4 {
			return 0
		}

		if nums[i] != nums[i+1] && nums[i+1] != nums[i+2] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]

}

func main() {

	array := []int{0, 1, 0, 1, 0, 1, 99}

	fmt.Println(singleNumber(array))

}
