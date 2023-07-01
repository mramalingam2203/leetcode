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
	for i := 0; i < len(s); i++ {
		if i != 0 && len(s)%i == 0 {
			fmt.Println(i, s[i])
		}
	}
	return true
}
