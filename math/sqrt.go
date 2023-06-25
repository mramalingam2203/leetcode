// https://leetcode.com/problems/sqrtx/

package main

import (
	"fmt"
	"math"
	_ "os"
	"strconv"
)

func babylonianGuess(x int) int {
	runesArray := []rune(strconv.Itoa(x))
	d := math.Ceil(float64(len(runesArray) / 2))
	fmt.Println(d)

	guess2 := 2 * math.Pow(10, d-1)
	guess7 := 7 * math.Pow(10, d-1)

	if math.Abs(math.Pow(float64(guess2), 2)-float64(x)) < math.Abs(math.Pow(float64(guess7), 2)-float64(x)) {
		return int(guess2)
	} else {
		return int(guess7)
	}

}

func mySqrt(x int) int {
	return 0
}

func main() {
	num := 20
	fmt.Println(babylonianGuess(num))
}
