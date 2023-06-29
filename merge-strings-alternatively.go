// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
)

func main() {

	s1 := "ab"
	s2 := "pqrs"

	fmt.Println(mergeAlternately(s1, s2))
}

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) < 1 || len(word1) < 1 || len(word2) > 100 || len(word1) > 100 {
		fmt.Println("words not valid")
		os.Exit(0)
	}

	runeS := []rune(word1)
	runeT := []rune(word2)
	_ = runeT
	fmt.Println(runeS)

	for i := range runeS {
		if runeS[i] < 97 || runeS[i] > 122 {
			fmt.Println("String 1 contanis invalid character")
			os.Exit(0)
		}
	}

	for i := range runeT {
		if runeT[i] < 97 || runeT[i] > 122 {
			fmt.Println("String 2 contanis invalid character")
			os.Exit(0)
		}
	}

	var result []rune

	if len(runeT) > len(runeS) {
		for i := 0; i < len(runeT); i++ {
			if i < len(runeS) {
				result = append(result, runeS[i])
			}
			result = append(result, runeT[i])
		}
	} else {
		for i := 0; i < len(runeS); i++ {
			result = append(result, runeS[i])
			if i < len(runeT) {
				result = append(result, runeT[i])
			}
		}
	}

	return string(result)
}
