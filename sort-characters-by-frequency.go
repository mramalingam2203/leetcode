// https://leetcode.com/problems/sort-characters-by-frequency/

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(frequencySort("heellooo"))
}

func frequencySort(s string) string {

	// Step 1: Create a map to count character frequencies
	charFrequency := make(map[rune]int)

	// Step 2: Populate the frequency map
	for _, char := range s {
		charFrequency[char]++
	}

	// Step 3: Create a custom sorting type and function
	type CharCount struct {
		Char  rune
		Count int
	}

	// Step 4: Sort the characters by frequency
	charCounts := []CharCount{}
	for char, count := range charFrequency {
		charCounts = append(charCounts, CharCount{char, count})
	}

	sort.Slice(charCounts, func(i, j int) bool {
		return charCounts[i].Count > charCounts[j].Count
	})

	// Step 5: Create a new string by appending characters in sorted order
	result := ""
	for _, cc := range charCounts {
		for i := 0; i < cc.Count; i++ {
			result += string(cc.Char)
		}
	}

	return result

}
