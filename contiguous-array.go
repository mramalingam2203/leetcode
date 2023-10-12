// https://leetcode.com/problems/contiguous-array/

package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 1, 1, 0, 1, 1}
	fmt.Println(generateSubarrays(array))
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
