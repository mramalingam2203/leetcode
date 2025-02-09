// https://leetcode.com/problems/find-the-difference/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	s1 := "ae"
	s2 := "aea"
	fmt.Println(findTheDifference(s1, s2))

}

func findTheDifference(s string, t string) byte {
	if len(s) < 0 || len(s) > 1e3 {
		fmt.Println("Strign length invalid")
		os.Exit(0)
	}

	if len(t) != len(s)+1 {
		fmt.Println("Strign lenghts not compatible")
		os.Exit(0)
	}

	runeS := []rune(s)
	runeT := []rune(t)
	fmt.Println(runeS)
	fmt.Println(runeT)

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

	// Sort the rune array 1
	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})

	// Sort the rune array 2
	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	for i := 0; i < len(runeT)-1; i++ {
		if runeT[i] != runeS[i] {
			return byte(runeT[i])
		}
	}

	return byte(runeT[len(runeT)-1])
}
