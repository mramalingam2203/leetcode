//https://leetcode.com/problems/happy-number/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(isHappy(19))
}

func isHappy(n int) bool {

	if n < 1 || n > math.MaxInt32 {
		return false
	}

	for n != 1 {
		s := strconv.Itoa(n)
		sum := 0
		for i := 0; i < len(s); i++ {
			temp, _ := strconv.Atoi(string(s[i]))
			sum += temp * temp
		}
		n = sum
	}

	return true
}
