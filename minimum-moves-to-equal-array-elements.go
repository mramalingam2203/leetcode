// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/

package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{1, 1, 1}
	// fmt.Println(arrayMin(array))
	fmt.Println(minMoves(array))
}

func minMoves(nums []int) int {

	minNum := arrayMin(nums)
	sum := arraySum(nums)

	// Calculate the minimum moves required
	result := sum - (minNum * len(nums))

	return result
}

func arraySum(arr []int) int {

	sum := 0

	for i := range arr {
		sum += arr[i]
	}

	return sum
}

func arrayMin(arr []int) int {

	min := math.MaxInt

	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
