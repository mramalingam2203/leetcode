// https://leetcode.com/problems/spiral-matrix-ii/

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(generateMatrix(1))

}

func generateMatrix(n int) [][]int {

	one := [][]int{{1}}
	two := [][]int{{1, 2}, {4, 3}}

	if n < 1 || n > 20 {
		fmt.Println("OUT OF RANGE!!!!")
		os.Exit(0)
	}

	if n == 1 {
		return one
	} else if n == 2 {
		return two
	}

	// Create a 2D slice (slice of slices)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 0
		}
	}

	num := 1
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1

	for num <= n*n {

		// Traverse top row
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}

		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Traverse bottom row
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}

		bottom--

		// Traverse left column
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix

}
