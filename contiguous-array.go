// https://leetcode.com/problems/contiguous-array/

package main

import "fmt"

func main() {
	array := []int{0, 1, 0, 1}
	fmt.Println(findMaxLength(array))
}

func findMaxLength(nums []int) int {
	subarrays := generateSubarrays(nums)
	fmt.Println(subarrays)
	result := 0
	var k int
	for subarray := range subarrays {
		k = isContiguous((subarrays[subarray]))
		if k > result && k != -1 {
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

	valueMap := make(map[int]int)

	for i := range arr {
		valueMap[arr[i]]++
	}

	fmt.Println(valueMap)

	if valueMap[0] == valueMap[1] {
		return len(arr)
	}

	return 0
}
