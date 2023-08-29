// https://leetcode.com/problems/fibonacci-number/

package main

import "fmt"

func main() {
	fmt.Println(fib(10))
}

func fib(n int) int {

	fib_table := make([]int, n+1)
	fib_table[0] = 0
	fib_table[1] = 1

	for i := 2; i < n+1; i++ {
		fib_table[i] = fib_table[i-1] + fib_table[i-2]
	}

	return fib_table[n]
}
