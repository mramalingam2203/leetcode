// https://leetcode.com/problems/contiguous-array/

package main

import "fmt"

func main() {
	array := []int{0, 1, 0}
	fmt.Println(findMaxLength(array))
}

func findMaxLength(nums []int) int {
	subarrays := generateSubarrays(nums)
	result := 0
	var k int
	for subarray := range subarrays {
		k = isContiguous((subarrays[subarray]))
		if k > result {
			result = k
		}
	}

	return result

}

func generateSubarrays(arr []int) [][]int {
	subarrays := [][]int{}

	for start := 0; start < len(arr); start++ {
		for end := start; end < len(arr); end++ {
			subarray := make([]int, end-start+1)
			copy(subarray, arr[start:end+1])
			subarrays = append(subarrays, subarray)
		}
	}

	return subarrays
}

func isContiguous(arr []int) int {

	count_0 := 0
	count_1 := 0

	for digit := range arr {
		if digit == 0 {
			count_0++
		} else {
			count_1++
		}
	}
	if count_0 == count_1 {
		return len(arr)
	}

	return -1
}
