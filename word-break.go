// https://leetcode.com/problems/word-break/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "leetycode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))

}

func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			for k := 0; k < len(wordDict); k++ {
				if dp[j] == true && strings.Contains(wordDict[k], s[j:i]) {

					dp[i] = true
					break
				}
			}

		}
	}

	return dp[n]

}
