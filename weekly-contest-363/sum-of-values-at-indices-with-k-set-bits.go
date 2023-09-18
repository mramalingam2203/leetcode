// https://leetcode.com/contest/weekly-contest-363/problems/sum-of-values-at-indices-with-k-set-bits/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	array := []int{5, 10, 1, 5, 2}
	target := 1
	sumIndicesWithKSetBits(array, target)

}

func sumIndicesWithKSetBits(nums []int, k int) int {

	for i := 0; i < len(nums); i++ {
		findSetBits(i, k)
	}

	return 0
}

func findSetBits(decimal int, k int) bool {

	if decimal == 0 {
		return false
	}

	binary := ""
	for decimal > 0 {
		// Use bitwise AND to get the least significant bit
		bit := decimal & 1
		// Convert the bit to a string and prepend it to the binary representation
		binary = strconv.Itoa(bit) + binary
		// Right shift the decimal number to check the next bit
		decimal >>= 1
	}

	if countOccurrences(strconv.Itoa(decimal), '1') == k {
		return true
	}
	return false
}

func countOccurrences(inputString string, targetChar byte) int {
	count := 0

	for i := 0; i < len(inputString); i++ {
		if inputString[i] == targetChar {
			count++
		}
	}
	fmt.Println(count)
	return count
}
