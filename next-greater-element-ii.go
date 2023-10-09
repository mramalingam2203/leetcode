// https://leetcode.com/problems/next-greater-element-ii/

package main

import "fmt"

func main() {

	array := []int{1, 2, 1}
	fmt.Println(nextGreaterElements(array))

}

func nextGreaterElements(nums []int) []int {

	result := []int{}
	for i, _ := range nums {
		result = append(result, circularSearch(nums, nums[i]))
	}

	return result

}

func circularSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid
		}

		// If the left half is sorted
		if arr[low] <= arr[mid] {
			if target >= arr[low] && target < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { // If the right half is sorted
			if target > arr[mid] && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1 // Element not found
}
