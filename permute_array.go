package main

import "fmt"

func main() {
	// Example input
	arr := []int{1, 2, 3}

	// Find permutations
	perms := permute(arr)

	// Print permutations
	for _, perm := range perms {
		fmt.Println(perm)
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result
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
