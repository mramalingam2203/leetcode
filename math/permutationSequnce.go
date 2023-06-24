//  https://leetcode.com/problems/permutation-sequence/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := 3
	k := 5
	fmt.Println(getPermutation(n, k))
}

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}

func getPermutation(n int, k int) string {

	// Constraints
	if n < 1 || n > 9 {
		fmt.Println("number out of range")
		os.Exit(0)
	}

	if k < 1 || k > factorial(n) {
		fmt.Println("inputs out of range")
		os.Exit(0)
	}
	// Find permutations
	// Example input
	arr := []int{}
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	// Find permutations
	perms := permute(arr)
	fmt.Println(perms[k-1])
	return convertIntToString(perms[k-1])
}

func convertIntToString(elems []int) string {
	b := ""
	for _, v := range elems {
		b += strconv.Itoa(v)
		b += ""
	}
	return b
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
