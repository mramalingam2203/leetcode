// https://leetcode.com/problems/find-peak-element/

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4}
	//findPeakElement(array)
	fmt.Println(max(array))
}

func findPeakElement(nums []int) int {

	n := len(nums)
	if n < 1 || n > 1000 {
		return 0
	}
	var peak int

	for i := 1; i < n-1; i++ {

		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			peak = i
		}

	}

	if n <= 3 {
		peak = max(nums)
	}

	return peak

}

func max(array []int) int {
	biggest := 0
	for i := 0; i < len(array); i++ {
		if array[i] > array[biggest] {
			biggest = i
		}
	}

	return biggest

}
