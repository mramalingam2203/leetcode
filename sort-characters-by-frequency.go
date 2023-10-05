// https://leetcode.com/problems/sort-characters-by-frequency/

package main

import "fmt"

func main() {

}

func frequencySort(s string) string {

	freq := make(map[byte]int)

	for _, char := range s {
		freq[char]++
	}

	fmt.Println(freq)

}
