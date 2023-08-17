// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

package main

import "fmt"

func search(nums []int, target int) bool {

	if len(nums) < 1 || len(nums) > 5000 || target < -1e4 || target > 1e4 {
		return false
	}

	if len(nums) == 1 && nums[0] == target {
		return true
	} else if len(nums) == 1 && nums[0] != target {
		return false
	}

	if len(nums) == 2 {
		if nums[0] == target || nums[1] == target {
			return true
		}
		return false
	}

	if len(nums) == 3 {
		if nums[0] == target || nums[1] == target || nums[2] == target {
			return true
		}
		return false
	}

	if len(nums) == 4 {
		return binarySearch(nums, target)
	}

	pivot := findPivotIndex(nums)

	if pivot < 0 {
		return true
	}
	fmt.Println("was here")

	if nums[pivot] == target {
		return true
	}
	fmt.Println("was here")

	if nums[pivot] > target {
		if binarySearch(nums[pivot:len(nums)-1], target) == true {
			return true
		}
	} else if nums[pivot] < target {
		if binarySearch(nums[0:pivot], target) == true {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func findPivotIndex(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		// Check if mid is the pivot index
		if mid > 0 && nums[mid] < nums[mid-1] {
			return mid
		}

		// Check if mid + 1 is the pivot index
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return mid + 1
		}

		// If the pivot is in the right half, search in the right
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1

}

func main() {

	array := []int{1, 1, 1, 3}
	target := 0
	//fmt.Println(array[findPivotIndex(array)])

	fmt.Println(search(array, target))

}
