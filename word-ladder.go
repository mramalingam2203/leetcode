// https://leetcode.com/problems/word-ladder/

package main

func main() {
	generatePossibleWords("hello")
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	//searc endword in wordlist -- if not found return zero

}

func generatePossibleWords(word string) {

	possibleWords := []string{}

	for i := 0; i < len(word); i++ {
		for char := 'a'; char <= 'z'; char++ {
			newWord := replaceChar(string(word), i, byte(char))
			if newWord != word {
				possibleWords = append(possibleWords, newWord)
			}

		}

	}

}

func replaceChar(input string, index int, newChar byte) string {
	// Convert the string to a byte slice
	bytes := []byte(input)

	// Update the byte at the specified index
	bytes[index] = newChar

	// Convert the byte slice back to a string
	return string(bytes)
}

/* set definitions*/
type Set map[string]bool

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item string) {
	s[item] = true
}

func (s Set) Contains(item string) bool {
	return s[item]
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Size() int {
	return len(s)
}
