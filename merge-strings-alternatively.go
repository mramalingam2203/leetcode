// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=programming-skills

package main

import "fmt"

func main() {

	s1 := "hi"
	s2 := "hello"

	fmt.Println(mergeAlternately(s1, s2))
}

func mergeAlternately(word1 string, word2 string) string {
	if len(word1) == 0 {
		return word2
	}


}
