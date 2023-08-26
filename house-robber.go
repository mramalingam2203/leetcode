// https://leetcode.com/problems/house-robber/

package main

func main() {

}

func rob(nums []int) int {

	n := len(nums)

	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2] + nums[i])

	return dp[n-1]

}
