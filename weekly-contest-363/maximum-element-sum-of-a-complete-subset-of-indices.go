// https://leetcode.com/contest/weekly-contest-363/problems/maximum-element-sum-of-a-complete-subset-of-indices/

package main

import "fmt"

func main() {
	array := []int{8, 7, 3, 5, 7, 2, 4, 9}
	//	maximumSum(array)
	subsets := findSubsets(array)
	findCompleteSubsets(subsets)
}

func findCompleteSubsets(array [][]int) {

	for i := 0; i < len(array); i++ {
		prod := 1
		for j := 0; j < len(array[i]); j++ {
			prod *= array[i][j]
		}
		if isPerfectSquare(prod) == true {
			fmt.Println(array[i], prod)
		}
	}

}

func maximumSum(nums []int) int64 {

	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPerfectSquare(nums[i]*nums[j]) == true {
				fmt.Println(nums[i], nums[j])
			}
		}
	}

	return 10
}

func findSubsets(arr []int) [][]int {
	var result [][]int

	var backtrack func(start int, currentSubset []int)
	backtrack = func(start int, currentSubset []int) {
		// Add the current subset to the result
		temp := make([]int, len(currentSubset))
		copy(temp, currentSubset)
		result = append(result, temp)

		// Explore subsets with the remaining elements
		for i := start; i < len(arr); i++ {
			currentSubset = append(currentSubset, arr[i])        // Include the current element
			backtrack(i+1, currentSubset)                        // Recursively find subsets
			currentSubset = currentSubset[:len(currentSubset)-1] // Exclude the current element
		}
	}

	backtrack(0, []int{})
	return result
}

func isPerfectSquare(n int) bool {

	// Base case: 0 and 1 are perfect squares
	if n <= 1 {
		return true
	}

	var left, right uint64 = 1, uint64(n)
	var mid, square uint64

	for left <= right {
		mid = left + (right-left)/2

		square = mid * mid

		if square == uint64(n) {
			return true
		} else if square < uint64(n) {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}
