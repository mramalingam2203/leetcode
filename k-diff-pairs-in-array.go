// https://leetcode.com/problems/k-diff-pairs-in-an-array/

package main

import "fmt"

func main() {

	array := []int{3, 1, 4, 1, 5}
	k := 2
	fmt.Println(findPairs(array, k))

}

func findPairs(nums []int, k int) int {
	result := []int{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if isKDiffPair(nums[i], nums[j], k) == true && nums[i] != nums[j] {
				result = append(result, nums[i], nums[j])
			}
		}
	}

	pairs := convertTo2D(result, len(result)/2, 2)
	count := 0
	temp := [2]int{}
	for pair, _ := range pairs {
		if pairs[pair][0] != temp[0] && pairs[pair][1] != temp[1] {
			count++
		}
		temp[0] = pairs[pair][0]
		temp[1] = pairs[pair][1]
	}
	return count
}

func convertTo2D(slice []int, rows, cols int) [][]int {
	if rows*cols != len(slice) {
		// The dimensions of the 2D slice do not match the length of the 1D slice.
		// You can handle this error in the way that's suitable for your application.
		return nil
	}

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
		copy(result[i], slice[i*cols:(i+1)*cols])
	}
	return result
}

func isKDiffPair(a, b, target int) bool {

	if abs(a-b) == target {
		return true
	}

	return false
}

func abs(a int) int {

	if a < 0 {
		return -1 * a
	}

	return a
}
