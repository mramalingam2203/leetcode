// https://leetcode.com/problems/word-ladder/

package main

func main() {
	generatePossibleWords("hello")
}

func generatePossibleWords(word string) {

	//possibleWords := []string{}

	for i := 0; i < len(word); i++ {
		for char := 'a'; char <= 'z'; char++ {
			newWord := replaceChar(string(word), i, byte(char))
			if newWord != word {

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
