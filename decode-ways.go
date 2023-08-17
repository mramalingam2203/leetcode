// https://leetcode.com/problems/decode-ways/

package main

func main() {

}

func numDecodings(s string) int {

	n := len(s)

	if n < 1 || n > 100{
		return 0
	}

	dp = make([]int, 0, n+1)

	dp[0] = 1

	if s[0] != '0'{
    	dp[1] = 1  
	else {
		dp[0] = 0
	}
	
	for i:=2; i <= n;i++ {
		oneDigit = strconv.parseInt(s[i - 1])
		twoDigits = strconv.parseInt(s[i - 2 : i])

		if oneDigit >= 1 and oneDigit <= 9{
			dp[i] = dp[i] + dp[i - 1]
		}

		if twoDigits >= 10 and twoDigits <= 26{
			dp[i] = dp[i] + dp[i - 2]
		}

	}
	
	return dp[n]
}
