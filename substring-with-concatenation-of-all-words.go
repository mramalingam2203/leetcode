// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

package main

import "fmt"

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	findSubstring(s, words)
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	wordCount := len(words)
	wordLength := len(words[0])
	totalWordsLength := wordCount * wordLength

	fmt.Println(wordCount, wordLength, totalWordsLength)

	if wordCount == 0 || totalWordsLength > len(s) {
		return result
	}

	wordFreqMap := make(map[string]int)

	for i := 0; i < wordLength-1; i++ {
		start := i
		currentWordCount := 0
		wordMap := wordFreqMap

		for j := i; j < len(s); j += wordLength {
			word = s[j : j+wordLength]
			if {

				currentWordCount++
				if currentWordCount == wordCount{
					resutl = append(resutl, start)
				}

			} else {

				currentWordCount = 0
				wordMap = wordFreqMap
				start += wordLength
				if start + totalWordsLength > len(s){
					break
				}

		}
	}

}

	return result
}
