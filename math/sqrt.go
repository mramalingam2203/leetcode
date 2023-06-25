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

	guess2 := 2 * math.Pow(10, d-1)
	guess7 := 7 * math.Pow(10, d-1)
	fmt.Println(guess2, guess7)
	if math.Abs(math.Pow(float64(guess2), 2)-float64(x)) < math.Abs(math.Pow(float64(guess7), 2)-float64(x)) {
		return int(guess2)
	} else {
		return int(guess7)
	}

}

func mySqrt(x int) int {
	var epsilon float64 = 1e-5
	var result, result_prev int
	result = babylonianGuess(x)

	for {
		result_prev = result
		result = (result_prev + x/result_prev) / 2
		if math.Abs(float64(result-result_prev)) > epsilon {
			break
		}
	}

	return result

}

func main() {
	num := 9981
	fmt.Println(mySqrt(num))
}
