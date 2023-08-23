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

	if n < 1 || n > 200 || len(triangle[0]) != 1 {
		return 0
	}

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {

		if i > 0 {
			if len(triangle[i]) != len(triangle[i-1])+1 {
				return 0
			}
		}

		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if triangle[i][j] < -1e4 || triangle[i][j] > 1e4 {
				return 0
			}
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
