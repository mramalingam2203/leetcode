package main

import (
	"fmt"
)

func maxArea(A []int) int {
	var area int
	length := len(A)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			area = max(area, min(A[j], A[i])*(j-i))
		}
	}
	return area
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	a := []int{1, 5, 4, 3}
	b := []int{3, 1, 2, 4, 5}
	fmt.Println(maxArea(a))
	fmt.Println(maxArea(b))
}
