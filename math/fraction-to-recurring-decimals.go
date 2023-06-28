// https://leetcode.com/problems/fraction-to-recurring-decimal/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if denominator == 0 || numerator < int(math.Pow(2.0, -31)) || denominator < int(math.Pow(2.0, -31)) || numerator > int(math.Pow(2.0, 31)) || denominator > int(math.Pow(2.0, 31)) {
		fmt.Println("Numbers out of range")
		os.Exit(0)
	}

	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	} else {
		//return fmt.Sprintf("%f", float64(numerator)/float64(denominator))
		findRecurrence(numerator, denominator)

	}

	return "hi"
}

func findRecurrence(num int, den int) {
	var res string
	quoient := num / den
	residuesMap := map[int]int{}
	remainder := num % den

	for{

 		if remainder !=0 && if residuesMap[remainder]
	}

}

func main() {
	num := 1
	den := 3
	fmt.Println(fractionToDecimal(num, den))
}
