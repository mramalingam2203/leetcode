// https://leetcode.com/problems/sum-of-squares-of-special-elements/

package main

import (
	"fmt"
	"os"
)

func sumOfSquares(nums []int) int {

	if len(nums) < 1 || len(nums) > 50 {
		fmt.Println("INvalid length")
		os.Exit(0)
	}

	sum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 50 {
			return 0
		}
		if len(nums)%(i+1) == 0 {
			sum += nums[i] * nums[i]
		}
	}
	return sum
}

func main() {

	list := []int{2, 7, 1, 19, 18, 3}
	sumOfSquares(list)

}
