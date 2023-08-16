// https://leetcode.com/problems/find-peak-element/

package main

import "fmt"

func main() {
	array := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(array))
}

func findPeakElement(nums []int) int {

	n := len(nums)
	if n < 1 || n > 1000 {
		return 0
	}

	slice := make([]int, len(nums)+2)

	slice[0] = nums[0] - 1
	slice[len(slice)-1] = nums[n-1] - 1

	for i := 0; i < len(nums); i++ {
		slice[i+1] = nums[i]
	}

	var peak int

	for i := 1; i < len(slice)-1; i++ {

		if slice[i] > slice[i-1] && slice[i] > slice[i+1] {
			peak = i - 1
		}

	}

	return peak

}
