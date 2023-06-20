package main

import (
	"fmt"
	"os"
)

func convert(s string, numRows int) string {
	if len(s) < 1 || len(s) > 1000 {
		fmt.Println("String length must be between 1 and 1000")
		os.Exit(0)
	}

	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	currentRow := 0
	direction := -1

	for _, char := range s {
		rows[currentRow] += string(char)
		if currentRow == 0 || currentRow == numRows-1 {
			direction *= -1
		}
		currentRow += direction
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}
