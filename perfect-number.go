// https://leetcode.com/problems/perfect-number/

package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(checkPerfectNumber(28))

}

func checkPerfectNumber(num int) bool {
	divisors := []int{}

	for i := 1; i <= int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
			if i != num/i {
				divisors = append(divisors, num/i)
			}
		}
	}

	sum := 0
	for i := range divisors {
		sum += divisors[i]
	}

	return sum-num == num

}
