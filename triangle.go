// https://leetcode.com/problems/triangle/

package main

func main() {
	array := [][]int{{1}, {1, 2, 3}, {3, 4, 5, 6}}
	minimumTotal(array)
}

func minimumTotal(triangle [][]int) int {

	n := len(triangle) // number of rows

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i <= n; i++ {
		dp[n-1][i] = triangle[n-1][i]
	}

	for i := n - 2; i > 0; i-- {

	}
}
