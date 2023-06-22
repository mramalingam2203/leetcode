package main

import (
	"fmt"
	"os"
)

func twoSum(nums []int, target int) []int {
	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109

	if len(nums) < 2 || len(nums) > 104 {
		fmt.Printf("Array undefined")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < -109 || nums[i] > 109 {
			fmt.Println("Numbers off range")
			os.Exit(0)
		}
	}

	if target < -109 || target > 109 {
		fmt.Printf("Target out of range")
		os.Exit(0)
	}

	

}
