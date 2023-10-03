// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

package main

import "fmt"

func main() {

	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	k := 8
	result := kthSmallest(matrix, k)
	fmt.Println(result) // Output: 13
}

func kthSmallest(matrix [][]int, k int) int {

	countLessEqual := func(target int) int {
		count := 0
		row, col := len(matrix)-1, 0

		for row >= 0 && col < len(matrix[0]) {
			if matrix[row][col] <= target {
				count += row + 1
				col++
			} else {
				row--
			}
		}

		return count
	}

	minVal, maxVal := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]

	for minVal < maxVal {
		mid := minVal + (maxVal-minVal)/2
		count := countLessEqual(mid)

		if count < k {
			minVal = mid + 1
		} else {
			maxVal = mid
		}
	}

	return minVal

}
