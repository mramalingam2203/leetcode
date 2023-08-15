// https://leetcode.com/problems/group-anagrams/

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(words))

}

func groupAnagrams(words []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result

}

func sortString(word string) string {

	letters := make([]string, len(word))
	s := ""
	for i := 0; i < len(word); i++ {
		letters[i] = string(word[i])
	}

	sort.Strings(letters)

	for i := 0; i < len(letters); i++ {
		s += letters[i]
	}

	return s

}
