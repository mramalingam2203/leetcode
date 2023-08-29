// https://leetcode.com/problems/base-7/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	convertToBase7(-7)
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func convertToBase7(num int) string {
	if num < -1e7 || num > 1e7 || num == 0 {
		return "0"
	}
	rems := []string{}
	var sign bool

	if num < 0 {
		num *= -1
		sign = true
	}
	for num != 0 {
		rems = append(rems, strconv.Itoa(num%7))
		num /= 7
	}

	result := reverse(strings.Join(rems, ""))

	if sign == true {
		result = "-" + result
	}

	fmt.Println(result)
	return result
}
