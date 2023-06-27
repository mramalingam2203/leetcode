// https://leetcode.com/problems/perfect-squares/

package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {

	if n < 1 || n > 1e4 {
		fmt.Println("Number out of range")
	}

	if math.Sqrt(float64(n))-math.Floor(math.Sqrt(float64(n))) == 0 {
		return 1
	}
	if n <= 3 {
		return n
	}
	res := n
	var temp int
	for x := 1; x <= n; x++ {
		temp = x * x
		if temp > n {
			break
		} else {
			res = min(res, 1+numSquares(n-temp))
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares(12))
}
