//https: //leetcode.com/problems/fraction-to-recurring-decimal/

package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {

	if numerator == 0 {
		return "0"
	}

	result := ""

	// Determine the sign of the result
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		result += "-"
	}

	// Calculate integer part
	result += strconv.Itoa(numerator / denominator)

	remainder := numerator % denominator
	if remainder == 0 {
		return result
	}

	result += "."

	remainMap := make(map[int]int)

	for remainder != 0 {

		if _, found := remainMap[remainder]; found {

			insertPosition := remainMap[remainder]
			recurringPart := result[insertPosition:]
			result = result[:insertPosition] + "(" + recurringPart + ")"
			break
		}

		// Add current remainder to the map
		remainMap[remainder] = len(result)

		remainder *= 10
		result += strconv.Itoa(remainder / denominator)
		remainder %= denominator

		//	fmt.Println(remainder)
	}

	return result
}

func abs(a int) int {

	if a < 0 {
		return -1 * a
	}

	return a

}

func main() {

	fmt.Println(fractionToDecimal(22, 7))
}
