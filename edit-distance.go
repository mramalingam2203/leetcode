// https://leetcode.com/problems/edit-distance/

package main

import "fmt"

func main() {
	s1 := "horse"
	s2 := "ros"
	minDistance(s1, s2)
	min(3, 2, 5)
}

func minDistance(word1 string, word2 string) int {

	m := len(word1)
	n := len(word2)

	if m < 0 || n < 0 || m > 500 || n > 500 {
		return 0
	}

	// Initialize a 2D array to store edit distances
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the first row and column
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Fill the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				fmt.Println(dp[i][j])
			} else {
				dp[i][j] = min(dp[i-1][j]+1, // Deletion
					dp[i][j-1]+1,   // Insertion
					dp[i-1][j-1]+1) // Substitution

			}

		}
	}

	return 0
}

func min(a, b, c int) int {
	if a < b {
		if b < c {
			return b
		} else {
			return c
		}
	}

	return a

}
