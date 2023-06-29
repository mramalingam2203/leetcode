// https://leetcode.com/problems/find-the-difference/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	s1 := "heleoworld"
	s2 := "helloworldy"
	fmt.Println(findTheDifference(s1, s2))

}

func findTheDifference(s string, t string) byte {
	if len(s) < 1 || len(s) > 1e3 {
		fmt.Println("Strign length invalid")
		os.Exit(0)
	}

	if len(t) != len(s)+1 {
		fmt.Println("Strign lenghts not compatible")
		return false
	}

	runeS := []rune(s)
	runeT := []rune(t)
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

}
