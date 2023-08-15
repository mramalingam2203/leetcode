// https://leetcode.com/problems/number-complement/

package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func findComplement(num int) int {
	s := intToBinary(num)
	fmt.Println(complement(s))

}

func complement(s str) int {

	for i:=0; i < len(s); i++ {
		


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
