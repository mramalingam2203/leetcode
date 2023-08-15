// https://leetcode.com/problems/substring-with-concatenation-of-all-words/

package main

func main() {
	s := "barfoothefoobarman"
	words = []string ["foo","bar"]
	findSubstring(s, words)
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	wordCount := len(words)
	wordLength := len(words[0])
	totalWordsLength := wordCount * wordLength

	if wordCount == 0 || totalWordsLength > len(s) {
		return result
	}

}
