// https://leetcode.com/problems/monotonic-array/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {
	nums := []int{6, 5, 5, 4}
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

	var inc, dec bool = true, true

	// Loop to check if array is increasing
	for i := 0; i < len(nums)-1; i++ {
		// To check if
		// array is not increasing
		if nums[i] > nums[i+1] {
			inc = false
		}
	}

	// Loop to check if array is decreasing
	for i := 0; i < len(nums)-1; i++ {

		// To check if
		// array is not decreasing
		if nums[i] < nums[i+1] {
			dec = false
		}
	}

	// Pick one whether inc or dec
	return inc || dec

}
