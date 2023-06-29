// https://leetcode.com/problems/sign-of-the-product-of-an-array/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {

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

	prod := 1

	for i := range nums {
		prod *= nums[i]
	}
}

func signFunc(n int) {

}
