//https://leetcode.com/problems/combination-sum-ii/

package main

import "fmt"

func generateCombinations(nums []int, k int) [][]int {
	var result [][]int
	backtrack(0, []int{}, nums, k, &result)
	return result
}

func backtrack(startIdx int, currentCombination []int, nums []int, k int, result *[][]int) {
	if len(currentCombination) == k {
		// Append a copy of the current combination to the result
		*result = append(*result, append([]int{}, currentCombination...))
		return
	}

	for i := startIdx; i < len(nums); i++ {
		currentCombination = append(currentCombination, nums[i])
		backtrack(i+1, currentCombination, nums, k, result)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}

func combinationSum2(candidates []int, target int) {
	var result [][]int
	for k := 0; k < len(candidates); k++ {
		result = append(result, generateCombinations(candidates, k))
	}
	fmt.Println(result)
}

func main() {
	nums := []int{1, 2, 3, 4}
	target := 13
	combinationSum2(nums, target)

}
