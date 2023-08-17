// https://leetcode.com/problems/edit-distance/


package main

func main() {
	s1 := "hello"
	s2 := "yellow"
}

func minDistance(word1 string, word2 string) int {

	if len(word1) < 0 || len(word2) < 0 { || len(word) > 500 || len(word2) > 500 {
		return 0
	}

	m := length(word1)
    n := length(word2)
    
    // Initialize a 2D array to store edit distances
    dp = make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}
    
    // Initialize the first row and column
    for i := 0 ; i <= m; i++ {
        dp[i][0] = i
	}

    for j := 0 ; j <= m; j++{
        dp[0][j] = j
	}

	fmt.Println(dp)
}