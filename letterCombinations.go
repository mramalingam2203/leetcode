package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	letters := map[int]string{2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}
	runeArray := []rune(digits)
	digi_letters := make([]string, 0)
	results := make([]string, 0)

	for i := 0; i < len(runeArray); i++ {
		digi_letters = append(digi_letters, letters[int(runeArray[i]-48)])
	}
	results = generateCombinations(digi_letters)
	return results

}

func generateCombinations(strings []string) []string {
	if len(strings) == 0 {
		return []string{""}
	}

	combinations := []string{}
	subCombinations := generateCombinations(strings[1:])

	for _, c := range subCombinations {
		for _, s := range strings[0] {
			combination := string(s) + c
			combinations = append(combinations, combination)
		}
	}

	return combinations
}

func main() {
	fmt.Println(letterCombinations("23"))
}
