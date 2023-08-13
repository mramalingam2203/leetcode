// https://leetcode.com/problems/single-number/

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

	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] < -3e4 || nums[i] > 3e4 {
			return 0
		}

		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]

}

func main() {

	array := []int{2, 2, 1, 1, 3, 3, 2, 4, 3, 3}

	fmt.Println(singleNumber(array))

}
