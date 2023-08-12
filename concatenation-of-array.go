// https: //leetcode.com/problems/concatenation-of-array/

package main

import (
	"fmt"
	"os"
)

func main() {

	list := []int{1, 2, 3, 4000, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(getConcatenation(list))

}

func getConcatenation(nums []int) []int {
	n := len(nums)
	i := 0

	if n < 1 || n > 1000 {
		fmt.Println("Invalid range")
		os.Exit(0)
	}
	for len(nums) < 2*n {
		if nums[i] < 1 || nums[i] > 1e3 {
			fmt.Println("Invalid range")
			os.Exit(0)
		}
		nums = append(nums, nums[i])
		i++

	}
	return nums

}
