// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

import "fmt"

func main() {
	array := []int{2, 2, 2, 0, 1}
	removeDuplicates(array)
	fmt.Println(findMin(array))
}

func removeDuplicates(nums []int) []int {
	uniqueMap := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			result = append(result, num)
		}
	}

	return result
}

func findMin(nums []int) int {
	if len(nums) < 1 || len(nums) > 5000 {
		return 0
	}
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			if nums[mid] < -5000 || nums[mid] > 5000 || nums[right] < -5000 || nums[mid] > 5000 {
				return 0
			}
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]

}
