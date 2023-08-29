// https://leetcode.com/problems/counting-bits/

package main

import "fmt"

func main() {

	fmt.Println(countBits(5))

}

func countBits(n int) []int {
	bitOnes := make([]int, n+1)
	for i := 0; i <= n; i++ {
		bitOnes[i] = counts(i)
	}

	return bitOnes
}

func counts(n int) int {
	count := 0
	for n != 0 {

		if n%2 == 1 {
			count++
		}

		n = n / 2
	}

	return count
}
