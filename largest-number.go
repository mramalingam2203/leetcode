// https://leetcode.com/problems/largest-number/

package main

import (
	"fmt"
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

	for i := 0; i < len(nums); i++ {

		if i != len(nums)-1 {
			fmt.Println(myCompare(s1[i], s1[i+1]))
		}
	}
}

func myCompare(str1, str2) bool {

	if str1+str2 > str2+str1 {
		return false
	}

	return true

}

func main() {

	array := []int{3, 30, 34, 5, 9}
	largestNumber(array)

}
