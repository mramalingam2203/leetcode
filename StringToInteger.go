package main

import (
	"fmt"
	_ "os"
)

const MIN int64 = -2147483648
const MAX int64 = 2147483647

func remove(items []string, item string) []string {
	newitems := []string{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}
func myAtoi(s string) int {

	fmt.Println(runeArray)
	return 0

}

func main() {
	myAtoi(" HELLO ")
}
