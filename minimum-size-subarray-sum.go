// https: //leetcode.com/problems/minimum-size-subarray-sum/

package main

import "math"

func main() {

	array := []int{2, 3, 1, 2, 4, 3}
	target := 7

}

func minSubArrayLen(target int, nums []int) int {

	n := len(nums)
	minLen = math.MaxInt32
	sum, left := 0, 0

	for right := 0; right < n-1; right++ {
		sum += nums[right]

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

}
