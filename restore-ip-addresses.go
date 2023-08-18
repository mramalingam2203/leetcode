// https://leetcode.com/problems/restore-ip-addresses/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	s := "25525511135"
	result := restoreIPAddresses(s)
	fmt.Println(result)

}

func restoreIPAddresses(s string) []string {

	var result []string
	var backtrack func(start int, current []string)

	backtrack = func(start int, current []string) {
		if start == len(s) && len(current) == 4 {
			result = append(result, strings.Join(current, "."))
			return
		}

		if len(current) >= 4 {
			return
		}

	}

	backtrack(0, nil)

	return result

}

func isValidSegment(s string) bool {

	if s[0] == '0' && len(s) > 1 {
		return false
	}

	num, _ := strconv.Atoi(s)
	if num >= 0 && num <= 255 {
		return true
	}

	return false

}
