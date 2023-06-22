package main

import (
	"fmt"
	"os"
	"regexp"
)

func myAtoi(s string) int {
	if len(s) < 0 || len(s) > 200 {
		fmt.Println("string length out of range")
		os.Exit(0)
	}

	// Regular expression to match alphabets
	regex := regexp.MustCompile("[a-zA-Z](?s)")

	// Replace all alphabets with an empty string
	s1 := regex.ReplaceAllString(s, "")

	runeArray := []rune(s1)

	fmt.Print(runeArray)
	return 0

}

func main() {
	s := " 01234 hello boy"
	myAtoi(s)
}
