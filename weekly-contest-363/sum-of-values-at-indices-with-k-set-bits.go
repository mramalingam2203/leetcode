// https://leetcode.com/contest/weekly-contest-363/problems/sum-of-values-at-indices-with-k-set-bits/

package main

import "strconv"

func main() {

}

func sumIndicesWithKSetBits(nums []int, k int) int {

}

func findSetBits(decimal, k int) bool {

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

	return true
}

func countOccurrences(inputString string, targetChar byte) int {
	count := 0

	for i := 0; i < len(inputString); i++ {
		if inputString[i] == targetChar {
			count++
		}
	}

	return count
}
