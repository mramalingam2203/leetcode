package main

import (
	"fmt"
)

func letterCombinations(digits string) {
	letters := map[int]string{2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}
	runeArray := []rune(digits)
	digi_letters := make([]string, 0)
	for i := 0; i < len(runeArray); i++ {
		digi_letters = append(digi_letters, letters[int(runeArray[i]-48)])
	}

}

func main() {
	letterCombinations("2357")
}
