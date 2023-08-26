// https://leetcode.com/problems/rotate-array/

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}

	rotate(array, 1)
	fmt.Println(array)

}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse(nums, 0, n-1) // Reverse the whole array
	reverse(nums, 0, k-1) // Reverse the first part up to k
	reverse(nums, k, n-1) // Reverse the second part from k to n-1
}

func reverse(nums []int, start, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}
