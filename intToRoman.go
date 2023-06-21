package main

import (
	"fmt"
	_ "os"
)

// Define the Roman numeral symbols and their corresponding values
var symbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Function to convert an integer to Roman numerals
func intToRoman(num int) string {
	roman := ""
	for _, s := range symbols {
		for num >= s.value {
			roman += s.symbol
			num -= s.value
		}
	}
	return roman
}

func main() {
	num := 1994
	roman := intToRoman(num)
	fmt.Printf("Roman numeral representation of %d is: %s\n", num, roman)
}
