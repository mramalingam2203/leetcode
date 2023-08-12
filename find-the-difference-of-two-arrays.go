// https://leetcode.com/problems/find-the-difference-of-two-arrays/

package main

import "fmt"

func findDifference(nums1 []int, nums2 []int) [][]int {

	diff := make([]int, 0, len(nums1))

	for i := 0; i < len(nums1); i++ {

		diff[i] = nums1[i] - nums2[i]

	}

	fmt.Println(diff)
}

func main() {
	list_1 := []int{1, 2, 3}
	list_2 := []int{2, 4, 6}

	findDifference(list_1, list_2)
}
