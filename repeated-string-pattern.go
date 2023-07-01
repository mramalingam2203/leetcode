// https://leetcode.com/problems/repeated-substring-pattern/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	//"os"
)

func main() {
	repeatedSubstringPattern("abcabcabcabcabc")
}

func repeatedSubstringPattern(s string) bool {
	runeS := []rune(s)

	runes_2D := make([][]rune, 3)
	runes_2D = convertTo2D(runeS, 3)
	fmt.Print(runes_2D)

	return true
}

func convertTo2D(slice []rune, columns int) [][]rune {
	length := len(slice)
	rows := length / columns
	if length%columns != 0 {
		rows++
	}

	result := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		start := i * columns
		end := start + columns
		if end > length {
			end = length
		}
		result[i] = slice[start:end]
	}

	return result
}
