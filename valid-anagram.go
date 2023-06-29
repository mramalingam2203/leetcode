// https://leetcode.com/problems/valid-anagram/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
	"sort"
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
	_ = runeT
	fmt.Println(runeS)

	for i := range runeS {
		if runeS[i] < 97 || runeS[i] > 122 {
			fmt.Println("String 1 contanis invalid character")
			return false
		}
	}

	for i := range runeT {
		if runeT[i] < 97 || runeT[i] > 122 {
			fmt.Println("String 2 contanis invalid character")
			return false
		}
	}

	// Sort the rune array 1
	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})

	// Sort the rune array 2
	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	for i := 0; i < len(runeS); i++ {
		if runeS[i] != runeT[i] {
			return false
		}
	}

	return true

}

func main() {
	s1 := "heleoworld"
	s2 := "dlrowolleh"
	fmt.Println(isAnagram(s1, s2))

}
