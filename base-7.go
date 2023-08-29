// https://leetcode.com/problems/base-7/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(convertToBase7(100))
}

func convertToBase7(num int) string {
	if num < -1e7 || num > 1e7 {
		return "0"
	}
	rems := []string{}
	for num != 0 {
		rems = append(rems, strconv.Itoa(num%7))
		num /= 7
	}

	return strings.Join(rems, "")
}
