package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   fly me   to   the moon     "
	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	// Split the string using space as the delimiter
	s1 := strings.Join(strings.Fields(s), " ")
	fmt.Println(s1)
	s2 := strings.Split(s1, " ")
	return len(s2[len(s2)-1])
	//return 9
}
