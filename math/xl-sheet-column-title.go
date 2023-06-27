// https://leetcode.com/problems/excel-sheet-column-title

package main

import (
	"fmt"
	"os"
)

func convertToTitle(columnNumber int) string {
	var s string
	var rem int

	for {
		rem = columnNumber % 26
		if rem == 0 {
			//	s = s + "Z")
			columnNumber = (columnNumber / 26) - 1
		} else {
			//s = s+ strconv.Itoa(rem-1)+"A")
			columnNumber /= 26"
		}
	}

	fmt.Print(s)
	return "hi"

}

func main() {
	//	convertToTitle(23)
}
