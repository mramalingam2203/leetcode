//  https://leetcode.com/problems/permutation-sequence/

package main

import (
	"fmt"
	_ "os"
)

func main() {
	n := 3
	k := 3
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	// Find permutations
	// Example input
	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	// Find permutations
	perms := permute(arr)
	_ = perms
	fmt.Println(perms)
	return "HI"
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
