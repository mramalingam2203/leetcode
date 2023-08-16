// https://leetcode.com/problems/maximum-product-subarray/

package main

import "fmt"

func main() {
	array := []int{-2, 0, -1}
	maxProduct(array)
}

func maxProduct(nums []int) int {

	if len(nums) < 1 || len(nums) > 2e4 {
		return 0
	}

	var subarrays [][]int
	var product, max int

	subarrays = generateSubarrays(nums)

	for i := 0; i < len(subarrays); i++ {
		product = 1
		for j := 0; j < len(subarrays[i]); j++ {
			product *= subarrays[i][j]
			if subarrays[i][j] < -10 || subarrays[i][j] > 10 {
				return 0
			}
		}
		if product > max {
			max = product
		}

	}
	fmt.Println(max)
	return 0
}

func generateSubarrays(arr []int) [][]int {
	var subarrays [][]int

	for start := 0; start < len(arr); start++ {
		for end := start + 1; end <= len(arr); end++ {
			subarray := []int{}
			for i := start; i < end; i++ {
				subarray = append(subarray, arr[i])
			}
			subarrays = append(subarrays, subarray)
		}
	}
	//fmt.Println(subarrays)
	return subarrays
}
