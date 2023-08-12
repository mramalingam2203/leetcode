// https://leetcode.com/problems/running-sum-of-1d-array/

package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4}
	fmt.Println(runningSum(list))
}

func runningSum(nums []int) []int {

	sums := make([]int, 0, len(nums))

	for _, i := range nums {
		if i != 0 {
			sums[i] = sums[i-1] + nums[i]
		} else {
			sums[i] = nums[i]
		}
	}

}
