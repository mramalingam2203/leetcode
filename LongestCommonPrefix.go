/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 || len(strs) > 200 {
		fmt.Println("String lenght out of bounds")
		os.Exit(0)
	}

	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < 0 || len(strs[i]) > 200 {
			fmt.Println("String lenght out of bounds")
			os.Exit(0)
		}
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	c := make([]string, 3)
	c[0] = "flower"
	c[1] = "fluid"
	c[2] = "fly"
	fmt.Println(longestCommonPrefix(c))

}
