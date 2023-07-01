// https://leetcode.com/problems/repeated-substring-pattern/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabcabc"))
}

func repeatedSubstringPattern(s string) bool {

	// Constraints

	if len(s) < 1 || len(s) > 1e4 {
		fmt.Println("repeated substring string invalid length")
		os.Exit(0)
	}

	runeS := []rune(s)
	for i := 0; i < len(runeS)/2; i++ {
		if i != 0 && len(runeS)%i == 0 {
			runeS_2D := make([][]rune, i)
			runeS_2D = convertTo2D(runeS, i)
			if allSubslicesEqual(runeS_2D) {
				return true
			}
		}
	}

	return false
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

func allSubslicesEqual(slice [][]rune) bool {
	if len(slice) == 0 {
		return true
	}

	firstSubslice := slice[0]

	for i := 1; i < len(slice); i++ {
		subslice := slice[i]
		if !equalSubslices(firstSubslice, subslice) {
			return false
		}
	}

	return true
}

func equalSubslices(subslice1, subslice2 []rune) bool {
	if len(subslice1) != len(subslice2) {
		return false
	}

	for i := 0; i < len(subslice1); i++ {
		if subslice1[i] != subslice2[i] {
			return false
		}
	}

	return true
}
