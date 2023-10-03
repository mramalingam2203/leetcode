// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

package main

import (
	"fmt"
	"sort"
)

func main() {

	arr_1 := []int{1, 7, 11}
	arr_2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(arr_1, arr_2, k))

}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			result = append(result, nums1[i])
			result = append(result, nums2[j])
		}
	}

	// Define the number of rows and columns for the 2D slice
	numRows := len(nums1) * len(nums2)
	numCols := 2

	// Create a 2D slice with the specified dimensions
	matrix := make([][]int, numRows)
	for i := range matrix {
		matrix[i] = make([]int, numCols)
	}

	// Populate the 2D slice with values from the 1D slice
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			// Calculate the index in the 1D slice corresponding to the 2D index
			index := i*numCols + j

			// Check if the index is within the bounds of the 1D slice
			if index < len(result) {
				matrix[i][j] = result[index]
			}
		}
	}

	// Define a custom sorting function based on the sum of elements
	sort.Slice(matrix, func(i, j int) bool {
		sum1 := sumOfElements(matrix[i])
		sum2 := sumOfElements(matrix[j])
		return sum1 < sum2
	})

	return matrix[0:k]
}

// Function to calculate the sum of elements in a slice
func sumOfElements(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}
