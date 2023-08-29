// https://leetcode.com/contest/weekly-contest-358/problems/max-pair-sum-in-an-array/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	array := []int{51, 71, 17, 24, 42}
	fmt.Println(maxSum(array))

}

func maxSum(nums []int) int {
	max := 0
	value := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			value = compareDigits(nums[i], nums[j])
			if value > max {
				max = value
			}
		}
	}

	if max == 0 {
		return -1
	}

	return max
}

func compareDigits(a, b int) int {

	str1, str2 := strconv.Itoa(a), strconv.Itoa(b)
	var max1, max2 int = 0, 0

	for i, _ := range str1 {
		temp, _ := strconv.Atoi(string(str1[i]))
		if temp > max1 {
			max1 = temp
		}
	}

	for i, _ := range str2 {
		temp, _ := strconv.Atoi(string(str2[i]))
		if temp > max2 {
			max2 = temp
		}
	}

	if max1 == max2 {
		return a + b
	}

	return 0

}
