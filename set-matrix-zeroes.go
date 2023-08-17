// https://leetcode.com/problems/set-matrix-zeroes/

package main

func main() {

	array := [][]int{{1, 1, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 1}, {1, 1, 1, 0}}
	setZeroes(array)

}

func setZeroes(matrix [][]int) {

	rows := len(matrix)
	cols := len(matrix[0])

	zeroRows := make(map[int]bool)
	zeroCols := make(map[int]bool)

	// Find the rows and columns with 0
	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroCols[j] = true
			}
		}
	}

	// Set entire rows to 0
	for row := range zeroRows {
		for j := 0; j < cols; j++ {
			matrix[row][j] = 0
		}
	}

	// Set entire columns to 0
	for col := range zeroCols {
		for i := 0; i < rows; i++ {
			matrix[i][col] = 0
		}
	}
}
