// https://leetcode.com/problems/find-the-difference-of-two-arrays/

package main

import (
	"fmt"
	"os"
)

func findDifference(nums1 []int, nums2 []int) [][]int {

	if len(nums1) < 1 || len(nums2) < 1 || len(nums1) > 1000 || len(nums2) > 1000 {
		fmt.Println("Invalid array lengths")
		os.Exit(0)
	}

	list_1 := []int{}
	list_2 := []int{}

	for i := 0; i < len(nums1); i++ {
		if nums1[i] < -1000 || nums1[i] > 1000 {
			fmt.Println("Invalid number values")
			os.Exit(0)
		}
		if !binarySearch(nums2, nums1[i]) {
			list_1 = append(list_1, nums1[i])
		}

	}

	for i := 0; i < len(nums2); i++ {
		if nums2[i] < -1000 || nums2[i] > 1000 {
			fmt.Println("Invalid number values")
			os.Exit(0)
		}

		if !binarySearch(nums1, nums2[i]) {
			list_2 = append(list_2, nums2[i])
		}
	}

	list_1 = removeDuplicates(list_1)
	list_2 = removeDuplicates(list_2)

	result := [][]int{list_1, list_2}
	fmt.Println(result)
	return result

}

func linearSearch(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func removeDuplicates(nums []int) []int {
	unique := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !unique[num] {
			unique[num] = true
			result = append(result, num)
		}
	}

	return result
}

func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	list_1 := []int{1, 2, 3, 3}
	list_2 := []int{1, 1, 2, 2}

	findDifference(list_1, list_2)
}
