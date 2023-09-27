// https://leetcode.com/problems/additive-number/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	return false
}

func main() {

	fmt.Println("12345")

}

func stringToNumberArray(input string) ([]int, error) {
	// Split the input string into individual number strings.
	numberStrings := strings.Fields(input)

	// Initialize a slice to hold the parsed integers.
	numbers := make([]int, len(numberStrings))

	// Parse each number string and store it in the numbers slice.
	for i, numStr := range numberStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// Return an error if parsing fails.
			return nil, err
		}
		numbers[i] = num
	}

	return numbers, nil
}
