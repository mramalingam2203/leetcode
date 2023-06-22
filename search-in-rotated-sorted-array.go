package main

import (
	"fmt"
	"os"
)

func search(nums []int, target int) int {
	if len(nums) < 1 || len(nums) > 5000 {
		fmt.Println("Array length out of range")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < -10000 || nums[i] > 10000 {
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

	if target < -10000 || target > 10000 {
		fmt.Println("target too low or high")
		os.Exit(0)
	}

	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1 // Target not found
}

func main() {
	// Example usage
	array := []int{4, 2, 9, 1, 7, 5}
	target := 10

	index := search(array, target)
	if index != -1 {
		fmt.Printf("Target found at index %d\n", index)
	} else {
		fmt.Println("Target not found in the array")
	}
}
