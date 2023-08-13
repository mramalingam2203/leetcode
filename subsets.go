/*
Spiral Matrix
Jump Game
Merge Intervals
Insert Interval
Spiral Matrix II
Unique Paths II
Minimum Path Sum
Set Matrix Zeroes
Search a 2D Matrix
Sort Colors
Subsets

*/

// https://leetcode.com/problems/subsets/

package main

import "fmt"

func subsets(nums []int) [][]int {
	allSubsets := [][]int{}
	generateSubsets(nums, 0, []int{}, &allSubsets)
	return allSubsets
}

func generateSubsets(array []int, index int, currentSubset []int, allSubsets *[][]int) {

	if index == len(array) {
		// Add the current subset to the list of allSubsets
		*allSubsets = append(*allSubsets, append([]int(nil), currentSubset...))
		return
	}

	// Exclude current element
	generateSubsets(array, index+1, currentSubset, allSubsets)
	fmt.Println(allSubsets)

	// Include current element
	currentSubset = append(currentSubset, array[index])
	generateSubsets(array, index+1, currentSubset, allSubsets)

	// Backtrack by removing the last element
	currentSubset = currentSubset[:len(currentSubset)-1]

}

func main() {
	list := []int{1, 2, 3}
	fmt.Println(subsets(list))
}
