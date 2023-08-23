// https://leetcode.com/problems/triangle/

package main

import "fmt"

func main() {
	array := [][]int{{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3}}
	fmt.Println(minimumTotal(array))
}

func minimumTotal(triangle [][]int) int {

	n := len(triangle) // number of rows

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {

			dp[i][j] += triangle[i][j] + min(dp[i+1][j], dp[i+1][j+1])

		}
	}

	return dp[0][0]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
