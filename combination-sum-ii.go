//https://leetcode.com/problems/combination-sum-ii/

package main

func generateCombinations(nums []int, k int) [][]int {
	var result [][]int
	backtrack(0, []int{}, nums, k, &result)
	return result
}

func combinationSum2(candidates []int, target int) [][]int {

}

func main() {

}
