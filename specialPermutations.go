package main

import (
	"fmt"
	"os"
)

func specialPerm(nums []int) int {

	if len(nums) < 2 || len(nums) > 14 {
		fmt.Println("Array length not good")
		os.Exit(0)
	}

	for num := range nums {
		if nums[num] < 1 || nums[num] > 109 {
			fmt.Println("Numbers out of range")
			os.Exit(0)
		}
	}

	return permute(nums)
}

func permute(nums []int) int {
	var result [][]int
	backtrack(nums, 0, &result)
	for idx := range result {
		fmt.Print(result[idx])
	}

	return 0
}

func backtrack(nums []int, start int, result *[][]int) {
	if start == len(nums) {
		// Found a permutation
		perm := make([]int, len(nums))
		copy(perm, nums)
		*result = append(*result, perm)
	} else {
		for i := start; i < len(nums); i++ {
			// Swap current element with the start element
			nums[start], nums[i] = nums[i], nums[start]

			// Recursively permute the remaining elements
			backtrack(nums, start+1, result)

			// Restore the original order for backtracking
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

func main() {
	a := make([]int, 3) // len(a)=5
	a[0] = 2
	a[1] = 3
	a[2] = 6

	fmt.Println(specialPerm(a))

}
