// https://leetcode.com/problems/4sum-ii/

package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {

	min := 0              // Minimum value for each number
	max := len(nums1) - 1 // Maximum value for each number
	length := 4           // Length of combinations

	var result [][]int
	current := make([]int, length)

	generateCombinations(min, max, length, current, &result)

	n := 0

	for i := 0; i < len(result); i++ {
		if nums1[result[i][0]]+nums2[result[i][1]]+nums3[result[i][2]]+nums4[result[i][3]] == 0 {
			n++
		}
	}

	return n

}

func generateCombinations(min, max, length int, current []int, result *[][]int) {
	if length == 0 {
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	for i := min; i <= max; i++ {
		current[len(current)-length] = i
		generateCombinations(min, max, length-1, current, result)
	}
}

func main() {
	nums1 := []int{0}
	nums2 := []int{0}
	nums3 := []int{0}
	nums4 := []int{0}

	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
