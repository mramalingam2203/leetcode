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
	fmt.Print(runeArray)
	return 0

}

func main() {
	s := " 4231"
	myAtoi(s)
}
