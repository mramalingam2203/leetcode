// https://leetcode.com/problems/restore-ip-addresses/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(isValidSegment("-255"))

}




func restoreIpAddresses(s) func() {
	
	var result []string
	var backtrack func(start int, current []string)

	return func  backtrack(start, current int){
		fmt.Println(result)

	}

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
