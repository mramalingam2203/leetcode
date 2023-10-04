package main

func findUniqueChars(input string) string {
	// Create a map to store the frequency of characters
	charCount := make(map[rune]int)
	uniqueChars := ""

	// Iterate through the input string
	for _, char := range input {
		// If the character has not been encountered before, add it to the result
		if charCount[char] == 0 {
			uniqueChars += string(char)
		}
		// Increment the character count in the map
		charCount[char]++
	}

	return uniqueChars
}

func intersection(str1, str2 string) string {
	// Create a map to store the frequency of characters in str1
	charCount := make(map[rune]int)
	commonChars := ""

	// Populate the charCount map with characters from str1
	for _, char := range str1 {
		charCount[char]++
	}

	// Iterate through str2 and check if each character exists in str1
	for _, char := range str2 {
		if charCount[char] > 0 {
			commonChars += string(char)
			charCount[char]--
		}
	}

	return commonChars
}

func canConstruct(ransomNote string, magazine string) bool {
	s := findUniqueChars(ransomNote)
	intersection(s, magazine)

}

func main() {
	s1 := "haha hi world how are you?"
	s2 := "haiworldeyu?"
	s3 := findUniqueChars(s1)
}
