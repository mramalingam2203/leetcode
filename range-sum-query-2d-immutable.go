// https://leetcode.com/problems/range-sum-query-2d-immutable/

package main

type NumMatrix struct {
	elements [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		elements: matrix,
	}

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for row := row1; row <= row2; row++ {
		for col := col1; col <= col2; col++ {
			sum += this.elements[row][col]
		}
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
