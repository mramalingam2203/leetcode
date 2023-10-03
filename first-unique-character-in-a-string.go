// https://leetcode.com/problems/first-unique-character-in-a-string/

package main

import "fmt"

func main() {

	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))

}

func firstUniqChar(s string) int {

	freq := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Iterate through the string to find the first unique character
	for i := 0; i < len(s); i++ {
		if freq[s[i]] == 1 {
			return i
		}
	}

	return -1
}
