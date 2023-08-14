// https://leetcode.com/problems/largest-number/

package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {

	s1 := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		s1[i] = strconv.Itoa(nums[i])
	}

	s2 := []string{}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[i]); j++ {
			s2 = append(s2, string(s1[i][j]))
		}
	}

	s3 := sort.Sort(sort.Reverse(sort.StringSlice(s2)))

	return "hi"
}

func main() {

	array := []int{3, 30, 34, 5, 9}
	largestNumber(array)

}
