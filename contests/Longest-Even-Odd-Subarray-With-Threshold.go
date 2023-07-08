// https://leetcode.com/contest/weekly-contest-352/problems/longest-even-odd-subarray-with-threshold/

package main

import (
	//"os"
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(
	findSubarrays(array)
	// for i := range array {
	// 	longestAlternatingSubarray(array[i], )
}

func longestAlternatingSubarray(nums []int, threshold int) int {

	return 0
}

func findSubarrays(arr []int) [][]int {
	n := len(arr)
	subarrays := [][]int{}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subarray := arr[i : j+1]
			subarrays = append(subarrays, subarray)

		}
	}

	return subarrays
}
