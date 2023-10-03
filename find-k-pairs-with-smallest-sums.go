// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

package main

import "fmt"

func main() {

	arr_1 := []int{1, 2, 3}
	arr_2 := []int{4, 5, 6}
	kSmallestPairs(arr_1, arr_2)

}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			fmt.Println(nums1[i], nums2[j])
		}
	}

}
