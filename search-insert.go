// https://leetcode.com/problems/search-insert-position/

package main

import (
	"fmt"
	"os"
)

func searchInsert(nums []int, target int) int {
	if len(nums) < 0 || len(nums) > 1e4 {
		fmt.Println("Array length out of range")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < -1e4 || nums[i] > 1e4 {
			fmt.Println("Number off the range")
			os.Exit(0)
		}
		temp := nums[i]
		if i < len(nums)-1 {
			if nums[i+1] >= temp {
				break
			}
		}
	}

	if target < -1e4 || target > 1e4 {
		fmt.Println("target too low or high")
		os.Exit(0)
	}

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		} else if nums[i-1] < target && nums[i+1] > target {
			return i
		}
	}

	return -1
}

func main() {
	// Example usage
	array := []int{1, 3, 5, 6}
	target := 9

	fmt.Println(searchInsert(array, target))
}
