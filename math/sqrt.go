// https://leetcode.com/problems/sqrtx/

package main

import (
	"fmt"
	"os"
)

func babylonianGuess(x int) {
	runesArray := []rune(x)
	d = math.Ceil(len(runesArray) / 2)
	guess2 = 2 * math.Pow(10, d-1)
	guess7 = 7 * math.Pow(10, d-1)
	if mag(math.Pow(guess2, 2)-x) < mag(math.Pow(guess7, 2)-x) {
		return guess2
	} else {
		return guess1
	}
}

func mySqrt(x int) int {
	return 0
}

func main() {
	num := 20
	babylonianGuess(num)
}
