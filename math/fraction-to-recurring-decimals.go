// https://leetcode.com/problems/fraction-to-recurring-decimal/

package main

import(
	_ "fmt"
	_ "os"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 || numerator < int(math.Pow(2.0, -31)) ||  denominator < int(math.Pow(2.0, -31))  ||  numerator > int(math.Pow(2.0, 31))  ||  denominator > int(math.Pow(2.0, 31)) {
		fmt.Println("Numbers out of range")
		os.Exit(0)
	}
    
}

func main)){

	fractionToDecimal(numerator int, denominator int)
}