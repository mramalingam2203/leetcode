package main

import (
	"fmt"
	//"os"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(diagonalSum(matrix))

}

func diagonalSum(mat [][]int) int {
	rows := len(mat)
	columns := len(mat)
	sum := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == j || i == rows-(j+1) {
				sum += mat[i][j]
			}
		}
	}
	/*
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if i == rows-(j+1) {
					//fmt.Println(i, j)
					fmt.Println(mat[i][j])
					sum += mat[i][j]
				}
			}
		}
	*/
	return sum
}
