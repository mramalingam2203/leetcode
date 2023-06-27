// https://leetcode.com/problems/armstrong-number

package main

import (
	_ "fmt"
	_ "os"
)

func power(x int, y uint64) int {

	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return power(x, y/2) * power(x, y/2)
	}

	return x * power(x, y/2) * power(x, y/2)
}

func main() {
	//    power(2,3)
}
