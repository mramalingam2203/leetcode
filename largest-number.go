// https://leetcode.com/problems/largest-number/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

	s2 := []string{}

	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[i]); j++ {
			s2 = append(s2, string(s1[i][j]))
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(s2)))

	s3 := strings.Join(s2, "")
	fmt.Println(s3)

	return s3
}

func main() {

	array := []int{3, 30, 34, 5, 9}
	largestNumber(array)

}
