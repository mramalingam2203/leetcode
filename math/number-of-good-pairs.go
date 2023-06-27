// https://leetcode.com/problems/number-of-good-pairs/

package main

import (
	"fmt"
	"os"
)

// Function to print combinations of size r
func printCombinations(arr []int, r int) {
	n := len(arr)
	if r > n {
		fmt.Println("Error: r is greater than the array size.")
		return
	}

	data := make([]int, r)
	combinationUtil(arr, n, r, 0, data, 0)
}

// Utility function to generate combinations recursively
func combinationUtil(arr []int, n, r, index int, data []int, dataIndex int) {
	if dataIndex == r {
		// Print the combination
		fmt.Println(data)
		return
	}

	if index >= n {
		return
	}

	data[dataIndex] = arr[index]
	combinationUtil(arr, n, r, index+1, data, dataIndex+1)
	combinationUtil(arr, n, r, index+1, data, dataIndex)
}

func numIdenticalPairs(nums []int) int {
	if len(nums) < 1 || len(nums) > 100 {
		fmt.Printf("Array length out of range")
		os.Exit(0)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 1 || nums[i] > 100 {
			fmt.Println("number out of range")
			os.Exit(0)
		}
	}
	return 0
}

func main() {
	// Example usage
	array := []int{1, 2, 3, 4}
	r := 2

	fmt.Printf("Combinations of size %d in the array %v:\n", r, array)
	printCombinations(array, r)
}
