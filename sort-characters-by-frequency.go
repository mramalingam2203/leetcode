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

	freq := make(map[byte]int)

	for _, char := range s {
		freq[byte(char)]++
	}

	// Create a slice to store key-value pairs
	var keyValuePairs []struct {
		Key   byte
		Value int
	}

	// Populate the slice with map data
	for key, value := range freq {
		keyValuePairs = append(keyValuePairs, struct {
			Key   byte
			Value int
		}{key, value})
	}

	// Define a custom sorting function based on values
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	result := ""
	// Print the sorted key-value pairs
	for _, kv := range keyValuePairs {
		result += string(kv.Key)
	}

	return result

}
