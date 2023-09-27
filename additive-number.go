// https://leetcode.com/problems/additive-number/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isAdditiveNumber("199100199"))

}

func isAdditiveNumber(num string) bool {
	// Create a slice to hold the characters.
	charArray := []rune{}

	// Iterate over the string and add each character to the slice.
	for _, char := range num {
		charArray = append(charArray, char)
	}

	numArray := []int{}

	for _, char := range num {
		number, _ := strconv.Atoi(string(char))
		numArray = append(numArray, number)
	}

	for i := 2; i < len(numArray); i++ {
		if numArray[i] != numArray[i-1]+numArray[i-2] {
			return false
		}
	}

	return true
}
