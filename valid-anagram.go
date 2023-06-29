// https://leetcode.com/problems/valid-anagram/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isAnagram(s string, t string) bool {
	if len(s) < 1 || len(t) < 1 || len(s) > 5e4 || len(t) > 5e4{
		fmt.Println("Strign lengths invalid")
		os.Exit(0)
	}

	if len(s) != len(t) {
		return false
		}
		

	for i := range s{
		if !unicode.IsLetter(s[i]){
			return false
		}
	}
}


func main{
	isAnagram("hello world", "world ellohe")

}