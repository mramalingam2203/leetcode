// https://leetcode.com/problems/minimum-moves-to-equal-array-elements/

package main

import "math"

func main() {}

func minMoves(nums []int) int {

	minNum = minimum value in the array nums
    sum = sum of all elements in the array nums

	# Calculate the minimum moves required
    minMoves = sum - (minNum * length of nums)
}


func arraySum(arr []int)int{

	sum := 0

	for i := range arr{
		sum += arr[i]
	}

	return sum
}



func arrayMin(arr []int)int{

	min := math.MinInt

	for i := range arr{
		if arr[i] < min{
			min = arr[i]
		}
	}
	return min
}