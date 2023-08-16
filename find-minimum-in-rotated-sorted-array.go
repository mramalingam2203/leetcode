// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

package main

func main() {
	array := []int{3, 4, 5, 1, 2}
	findMin(array)
}

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	for {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}

		if left > right {
			break
		}
	}

	return nums[left]

}
