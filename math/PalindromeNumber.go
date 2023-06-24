package main

import (
	"fmt"
	"math"
	_ "os"
	"strconv"
)

func isPalindrome(x int) bool {
	digitArray := strconv.Itoa(x)
	//fmt.Println(digitArray[0] - 48)
	max_tens := math.Pow(10, float64(len(digitArray)-1))

	for i := 0; i < len(digitArray)-1; i++ {
		if x%int(max_tens) == int(digitArray[i+len(digitArray)-1])-48 {
			fmt.Println(int(digitArray[i+len(digitArray)-2]) - 48)
			max_tens /= 10
		}
	}

	return true

}

/*
func isPalindrome0(x int) bool {
	numstring := strconv.Itoa(x)
	runeArray := []rune(numstring)
	reverseSlice(runeArray)
	if numstring == runeArray {
		return true
	}
	return false

}

func reverseSlice(slice []rune) {
	length := len(slice)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
}
*/

// rune values 1 through 9 ---> 49 through 57
func main() {
	var num int = 123321

	fmt.Println(isPalindrome(num))
}
