// https://leetcode.com/problems/word-break/

package main

import (
	"fmt"
)

func main() {
	s := "catsandog"
	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, wordDict))

}

func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	dp := make([]bool, n+1)
	wordSet := make(map[string]bool)

	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp[0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k < len(wordDict); k++ {
				if dp[j] == true && wordSet[s[j:i]] {

					dp[i] = true
					break
				}
			}

		}
	}

	return dp[n]

}
