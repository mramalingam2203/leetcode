// https://leetcode.com/problems/monotonic-array/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(isMonotonic(nums))
}

func isMonotonic(nums []int) bool {

	//constraints
	if len(nums) < 1 || len(nums) > 1e5 {
		fmt.Println("integer array length out of range")
		os.Exit(0)
	}

	for i := range nums {
		if nums[i] < -1e5 || nums[i] > 1e5 {
			fmt.Println("integer array elements invalid")
			os.Exit(0)
		}
	}

	var result bool

	for i := 0; i < nums.Len(); i++ {
		if nums[i] < nums[i+1] {
			result &= true
		} else {
			result &= false
		}
	}

	return result
}
