package main

import (
	"fmt"
	"os"
	"strings"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < 1 || len(needle) < 1 || len(haystack) > 1e4 || len(needle) > 1e4 {
		fmt.Println("Invalid strings")
		os.Exit(0)
	}

	return strings.Index(haystack, needle)
}

func main() {
	strStr("leet", "leeto")
}
