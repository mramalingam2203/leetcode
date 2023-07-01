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
	fmt.Println(runeS)
	return true
}
