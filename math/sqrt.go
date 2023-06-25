// https://leetcode.com/problems/sqrtx/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func babylonianGuess(x int) int {

	runesArray := []rune(strconv.Itoa(x))
	d := math.Ceil(float64(len(runesArray) / 2))
	fmt.Println(d)

	if d == 0 {
		if x >= 1 && x <= 3 {
			return 1
		} else if x >= 4 && x < 9 {
			return 2
		} else if x == 9 {
			return 3
		}
	}

	guess2 := 2 * math.Pow(10, d-1)
	guess7 := 7 * math.Pow(10, d-1)
	fmt.Println(guess2, guess7)
	if math.Abs(math.Pow(float64(guess2), 2)-float64(x)) < math.Abs(math.Pow(float64(guess7), 2)-float64(x)) {
		return int(guess2)
	} else {
		return int(guess7)
	}

}

func smaller(a int, b int) int {
	if a < b {
		return a
	}
	return b

}
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	var epsilon float64 = 1e-9
	var result, result_prev int
	result = babylonianGuess(x)
	count := 0
	for {
		result_prev = result
		result = (result_prev + x/result_prev) / 2
		if math.Abs(float64(result-result_prev)) < epsilon || count > 1e5 {
			return smaller(result, result_prev)
		}
		count++
		fmt.Println(count, result)

	}

	return result

}

func main() {
	num := 8
	fmt.Println(mySqrt(num))
}
