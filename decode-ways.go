// https://leetcode.com/problems/decode-ways/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(numDecodings("06"))

}

func numDecodings(s string) int {

	n := len(s)

	if n < 1 || n > 100 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 1

	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[0] = 0
	}

	for i := 2; i < n+1; i++ {

		oneDigit, _ := strconv.ParseInt(string(s[i-1]), 10, 32)
		twoDigits, _ := strconv.ParseInt(string(s[i-2:i]), 10, 32)

		if oneDigit >= 1 && oneDigit <= 9 {
			dp[i] = dp[i] + dp[i-1]
		}

		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] = dp[i] + dp[i-2]
		}

	}

	return dp[n]
}
