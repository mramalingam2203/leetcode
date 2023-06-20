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
    
	vals := []rune{ 0, 1,2, 3, 4, 5, 6, 7,8, 9,+, -}

	for i := 0; i < len(runeArray); i++ {
		if runeArray[i] < 48 || runeArray[i] > 57 {
			runeArray = append(runeArray[:i], runeArray[i+1:]...)
			//fmt.Println(runeArray[i])
		}

	}

	fmt.Print(runeArray)
	return 0

}

func main() {
	s := " 01234 h"
	myAtoi(s)
}
