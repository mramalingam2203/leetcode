package main

import (
	"fmt"
	"os"
)

func searchRange(nums []int, target int) []int {

	if len(nums) < 0 || len(nums) > 100000 {
		fmt.Println("Array length out of range")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < -1e9 || nums[i] > 1e9 {
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

	if target < -1e9 || target > 1e9 {
		fmt.Println("target too low or high")
		os.Exit(0)
	}

	var indices []int

	for i, num := range nums {
		if num == target {
			fmt.Println(i)
			indices = append(indices, i)
		}
	}
	var ret_idx []int

	if len(indices) == 0 {
		ret_idx = append(ret_idx, -1)
		ret_idx = append(ret_idx, -1)
		return ret_idx

	}

	ret_idx = append(ret_idx, indices[0])
	ret_idx = append(ret_idx, indices[len(indices)-1])

	return ret_idx

}

func main() {
	// Example usage
	array := []int{5, 7, 7, 8, 8, 10}
	target := 6

	index := searchRange(array, target)
	fmt.Print(index)
}
