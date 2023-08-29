// https://leetcode.com/problems/n-th-tribonacci-number

package main

import "fmt"

func main() {
	fmt.Println(tribonacci(2))
}

func tribonacci(n int) int {

	if n <= 0 || n > 37 {
		return 0
	}

	trib_table := make([]int, n+1)
	trib_table[0] = 0
	trib_table[1] = 1
	trib_table[2] = 1

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}

	for i := 3; i <= n; i++ {
		trib_table[i] = trib_table[i-1] + trib_table[i-2] + trib_table[i-3]
	}

	return trib_table[n]
}
