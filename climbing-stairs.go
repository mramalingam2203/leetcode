// https://leetcode.com/problems/climbing-stairs/

package main

import "fmt"

func main() {
	n := 2
	fmt.Println(climbStairs(n))

}

func climbStairs(n int) int {

	if n <= 2 {
		return 0
	}

	ways := make([]int, n+1)

	ways[0] = 1
	ways[1] = 2

	// Calculate the number of ways for each step using dynamic programming
	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]

}
