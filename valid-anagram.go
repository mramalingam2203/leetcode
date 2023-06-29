// https://leetcode.com/problems/valid-anagram/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
	"unicode"
)

func isAnagram(s string, t string) bool {
	if len(s) < 1 || len(t) < 1 || len(s) > 5e4 || len(t) > 5e4 {
		fmt.Println("Strign lengths invalid")
		os.Exit(0)
	}

	if len(s) != len(t) {
		fmt.Println("Strign unequal lenghts")
		return false
	}
	runeS := []rune(s)
	runeT := []rune(t)

	for i := range runeS {
		if !unicode.IsLetter(runeS[i]) {
			fmt.Println("String 1 contanis invalid character")
			return false
		}
	}

	for i := range runeT {
		if !unicode.IsLetter(runeT[i]) {
			fmt.Println("String 2 contanis invalid character")
			return false
		}
	}

	return true
}

func main() {
	s1 := "hello world"
	s2 := "world el1lo"
	fmt.Println(isAnagram(s1, s2))

}
