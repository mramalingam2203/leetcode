// https://leetcode.com/problems/sign-of-the-product-of-an-array/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {
	array := []int{9, 72, 34, 29, 0, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(array))
}

func arraySign(nums []int) int {
	if len(nums) < 1 || len(nums) > 1e3 {
		fmt.Println("array length invalid")
		os.Exit(0)
	}

	for i := range nums {
		if nums[i] < -100 || nums[i] > 100 {
			fmt.Println("Invalid numbers")
			os.Exit(0)
		}
	}

	return signFunc(nums)
}

func signFunc(nums []int) int {
	var negative int
	for i := range nums {
		if nums[i] < 0 {
			negative++
		} else if nums[i] == 0 {
			return 0
		}
	}
	fmt.Println(negative)
	if negative%2 != 0 {
		return -1
	}

	return 1

}
