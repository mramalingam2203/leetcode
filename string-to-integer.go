package main

import (
	"fmt"
	"os"
)

func myAtoi(s string) int {
	if len(s) < 0 || len(s) > 200 {
		fmt.Println("string length out of range")
		os.Exit(0)
	}
	runeArray := []rune(s)

	for i := 0; i < len(runeArray); i++ {
		if runeArray[i] < 48 || runeArray[i] > 57 {
			runeArray = append(runeArray[:i], runeArray[i+1:]...)

		}

	}

	fmt.Print(runeArray)
	return 0

}

func main() {
	s := " 0123456789 hello boy"
	myAtoi(s)
}
