// https://leetcode.com/problems/largest-number/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	if len(nums) < 1 || len(nums) > 100 {
		return "0"
	}

	s1 := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 1e9 {
			return "0"
		}

		s1[i] = strconv.Itoa(nums[i])
	}

	// Sort the slice using the compare function
	sort.Slice(nums, func(i, j int) bool {
		return myCompare(s1[i], s1[j])
	})

	// Sort the slice in descending order using a custom comparator
	sort.Slice(s1, func(i, j string) bool {
		return s1[i] > s2[j] // Compare in descending order
	})

	fmt.Println(s1)

	return "hi"
}

func myCompare(str1 string, str2 string) bool {

	fmt.Println(str1+str2, str2+str1)
	if str1+str2 > str2+str1 {
		return false
	}

	return true

}

func main() {

	array := []int{3, 30, 34, 5, 9}
	largestNumber(array)

}
