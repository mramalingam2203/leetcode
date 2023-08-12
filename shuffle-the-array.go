// https://leetcode.com/problems/shuffle-the-array/

package main

import (
	"fmt"
	"os"
)

func main() {

	Nums := []int{2, 5, 1, 3, 4, 7}
	N := 2
	fmt.Println(shuffle(Nums, N))

}

func shuffle(nums []int, n int) []int {

	if n < 1 || n > 500 {
		fmt.Println("Invalid number")
		os.Exit(0)
	}

	for i := 1; i < len(nums)-n; i++ {
		nums[i], nums[i+n] = nums[i+n], nums[i]
	}

	fmt.Println(nums)

	return nums

}
