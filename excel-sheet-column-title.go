// https://leetcode.com/problems/excel-sheet-column-title/

package main

import "math"

func convertToTitle(columnNumber int) string {

	if columnNumber < 1 || columnNumber > math.MaxInt32 {
		return "Invalid"
	}
	title := ""
	var reminder int
	for columnNumber > 0 {
		reminder = (columnNumber - 1) / 26
		title += string(rune(reminder + 65))
		columnNumber = (columnNumber - 1) / 26
	}
	runes := []rune(title)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	convertToTitle(28)
}
