// https://leetcode.com/problems/longest-uncommon-subsequence-ii/

package main

import "fmt"

func main() {

	stringArray := []string{"aaa", "acb"}
	fmt.Println(findLUSlength(stringArray))

}

func findLUSlength(strs []string) int {

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(strs[i-1]) {
			strs[i], strs[i-1] = strs[i-1], strs[i]
		}
	}

	for i := 1; i < len(strs); i++ {
		for j := i - 1; j > 0; j-- {
			if is_subsequence(strs[j], strs[i]) == false {
				return len(strs[i])
			}
		}
	}

	return -1
}

func is_subsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[i] {
			i++
		}
		j++
	}
	return i == len(s)
}
