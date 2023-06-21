package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		for j := 1; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue
			}
			return i
		}
	}
	return -1

}

func main() {
	fmt.Println(strStr("leetcode", "lee"))
}
