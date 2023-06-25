// https://leetcode.com/problems/plus-one/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func convertIntToString(elems []int) string {
	b := ""
	for _, v := range elems {
		b += strconv.Itoa(v)
		b += ";"
	}
	return b
}

func plusOne(digits []int) []int {
	if len(digits) == 0  || len(digits) > 100{
		fmt.Println("int array lenngth out of range")
		os.Exit(0)
	}
	
	for i range len(digits) {
		for j range digits[i] {
			if digits[i][j] < 1 || digits[i][j] >9{
				fmt.Println("numbers out of range")
				os.Exit(0)
			}
		}
	}

	if digit[0] == '0' {
		fmt.Println("number invalid")
		os.Exit(0)
	}

	fmt.Println(convertIntToString(digits))
	return digits
}

func main() {
	a := []int{1, 2, 3}
	plusOne(a)
}
