// https://leetcode.com/problems/number-complement/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(findComplement(1))
}

func findComplement(num int) int {
	s := intToBinary(num)
	return complement(s)
}

func complement(s string) int {

	empty := ""

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			empty += "1"
		} else {
			empty += "0"
		}
	}

	intValue, err := strconv.ParseInt(empty, 2, 0)
	_ = err
	return int(intValue)

}

func intToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binaryStr := ""

	for n > 0 {
		// Get the remainder when dividing by 2
		remainder := n % 2
		// Prepend the remainder to the binary string
		binaryStr = strconv.Itoa(remainder) + binaryStr
		// Divide the number by 2
		n /= 2
	}

	return binaryStr
}
