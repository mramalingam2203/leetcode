//https://leetcode.com/problems/roman-to-integer/

/*
   - Split the Roman Numeral string into Roman Symbols (character).
   - Convert each symbol of Roman Numerals into the value it represents.
   - Take symbol one by one from starting from index 0:
       If current value of symbol is greater than or equal to the value of next symbol, then add this value to the running total.
       else subtract this value by adding the value of next symbol to the running total.
*/

package main

import (
	"fmt"
	"os"
)

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: romanToInt romanNumeral")
		os.Exit(1)
	}

	romanNumeral := os.Args[1]
	romanToInt(romanNumeral)

	//	var total int = 0
}

func romanToInt(s string) int {
	sum := 0
	rtoIMap := make(map[int]int)
	rtoIMap[73] = 1    //I
	rtoIMap[86] = 5    //V
	rtoIMap[88] = 10   //X
	rtoIMap[76] = 50   //L
	rtoIMap[67] = 100  //C
	rtoIMap[68] = 500  //D
	rtoIMap[77] = 1000 //M

	romanSymbols := []rune(s)
	for i := 0; i < len(romanSymbols); i++ {
		if i+1 < len(romanSymbols) {
			fmt.Println(rtoIMap[int(romanSymbols[i])])
			if rtoIMap[int(romanSymbols[i])] >= rtoIMap[int(romanSymbols[i+1])] {
				sum += rtoIMap[int(romanSymbols[i])]
			} else {
				sum = sum + rtoIMap[int(romanSymbols[i+1])] - rtoIMap[int(romanSymbols[i])]
				i++
			}
		} else {
			sum += rtoIMap[int(romanSymbols[i])]
		}

	}

	fmt.Println(sum)
	return sum
}
