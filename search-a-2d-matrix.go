//https://leetcode.com/problems/search-a-2d-matrix/

package main

import "fmt"

func main() {

	array := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 11
	fmt.Println(searchMatrix(array, target))
}

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	left := 0
	right := m*n - 1

	for left <= right {

		mid := left + (right-left)/2
		midValue := matrix[mid/n][mid%n]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
