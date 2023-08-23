// https://leetcode.com/problems/pascals-triangle-ii/

package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(generate(5))

}

func getRow(rowIndex int) []int {
	return generate(rowIndex)

}

func generate(numRows int) []int {

	if numRows == 0 {
		fmt.Println("No rows")
		os.Exit(0)
	}

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {

		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle[numRows-1]
}
