// https://leetcode.com/problems/longest-increasing-subsequence/

package main

import "fmt"

func main() {

	array := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLIS(array))

}

func lengthOfLIS(nums []int) int {

	n := len(nums)
	dp := make([]int, n)

	maxLen := 1

	for i := 0; i < n; i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// Update the maximum LIS length
		maxLen = max(maxLen, dp[i])
	}

	return maxLen

}

func max(a, b int) int {

	if a < b {
		return b
	}

	return a

}
