// https://leetcode.com/problems/climbing-stairs/

package main

import "fmt"

func main() {
	n := 3
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {

	if n < 1 || n > 45 {
		return 0
	}

	if n <= 2 {
		return n
	}

	ways := make([]int, n+1)

	ways[0] = 1
	ways[1] = 2

	// Calculate the number of ways for each step using dynamic programming
	for i := 2; i < n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n-1]

}
