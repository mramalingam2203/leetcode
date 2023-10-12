// https://leetcode.com/problems/reverse-string-ii/

package main

import "fmt"

func main() {

	s := "abcdefg"
	k := 2
	//reverseStr(s, k)
	fmt.Println(reverse(s))

}

/*func reverseStr(s string, k int) string {

	for i := 0; i < len(s); i++ {
		if

}
*/

func reverse(s string) string {
	// Convert the string to a rune slice.
	runes := []rune(s)

	// Reverse the order of the runes.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the reversed rune slice back to a string.
	reversedString := string(runes)

	return reversedString
}
