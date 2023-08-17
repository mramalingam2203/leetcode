// https://leetcode.com/problems/spiral-matrix-ii/

package main

func main() {

}

func generateMatrix(n int) [][]int {

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
		for i :=left; i <= right; i++{
			matrix[top][i] = num
			num = num++
		}

		top++


        // Traverse right column
        for i := top; i <= bottom; i++){
            matrix[i][right] = num
            num = num++
		}
        right--

	}

}
