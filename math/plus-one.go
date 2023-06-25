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
	fmt.Println(convertIntToString(digits))

}

func main() {
	a := []int{1, 2, 3}
	plusOne(a)
}
