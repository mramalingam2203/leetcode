// https://leetcode.com/problems/word-break/

package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= n; i++ {
		for j:=0; j <= i; j++{
			if dp[j] == true && s[j:i] {

				dp[i] = true
				break
			}

		}
	return dp[n]
}
