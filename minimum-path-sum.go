// https://leetcode.com/problems/minimum-path-sum/

package main

import "fmt"

func minPathSum(grid [][]int) int {

	m := len(grid)
	n := len(grid[0])

	if m < 1 || m > 200 || n < 1 || n > 200 {
		return 0
	}

	// Create a DP table of the same size as the grid
	// Define the size of the 2D slice

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize the DP table with the value of the first cell
	dp[0][0] = grid[0][0]

	// Fill in the DP table
	for i := 1; i <= m-1; i++ {

		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j <= n-1; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i <= m-1; i++ {
		for j := 1; j <= n-1; j++ {

			if grid[i][j] < 0 || grid[i][j] > 200 {
				return 0
			}

			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]

}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {

	mesh := [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(minPathSum(mesh))

}
