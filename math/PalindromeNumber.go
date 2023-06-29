package main

import (
	"fmt"
	//"math"
	_ "os"
	"strconv"
)

func isPalindrome(x int) bool {
	numstring := strconv.Itoa(x)
	runeArray := []rune(numstring)
	tmpArray := []rune(numstring)

	reverseSlice(runeArray)

	for i := range tmpArray {
		if runeArray[i] != tmpArray[i] {
			return false
		}
	}
	return true

}

func reverseSlice(slice []rune) {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// rune values 1 through 9 ---> 49 through 57
func main() {
	var num int = 123321

	fmt.Println(isPalindrome(num))
}
