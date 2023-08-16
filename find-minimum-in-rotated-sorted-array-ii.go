// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

import "fmt"

func main() {
	array := []int{3, 3, 1, 3}
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
	uniques := removeDuplicates(nums)

	if len(uniques) < 1 || len(uniques) > 5000 {
		return 0
	}
	left := 0
	right := len(uniques) - 1

	for left < right {
		mid := (left + right) / 2
		if uniques[mid] > uniques[right] {
			if uniques[mid] < -5000 || uniques[mid] > 5000 || uniques[right] < -5000 || uniques[mid] > 5000 {
				return 0
			}
			left = mid + 1
		} else {
			right = mid
		}
	}

	return uniques[left]

}
