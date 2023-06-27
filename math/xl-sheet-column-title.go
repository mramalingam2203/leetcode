// https://leetcode.com/problems/excel-sheet-column-title

package main

import (
	"fmt"
	_ "os"
)

/*
func convertToTitle(columnNumber int) string {
	var s string
	var rem int

	for {
		rem = columnNumber % 26
		if rem == 0 {
			//	s = s + "Z")
			columnNumber = (columnNumber / 26) - 1
			fmt.Println(columnNumber)
		} else {
			//s = s+ strconv.Itoa(rem-1)+"A")
			columnNumber /= 26
		}
	}

	fmt.Print(s)
	return "hi"

}
*/

func convertToTitle(num int) string {
	var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var res string = ""

	if num < 26 {
		res = res +
		fmt.Println(alpha[num-1])
		return res
	} else{
		q := (num / 26)
		r := num % 26
		res := ""

		if(r == 0){
		  if(q == 1){
			res.convertToTitle(1,alpha[(26 + r-1)%26]);
		  }
		  else{
			res = num_hash(q-1);
			res.append(1,alpha[(26 + r-1)%26]);
		  }
		}
		else{
		  res = num_hash(q);
		  res.append(1,alpha[(26 + r-1)%26]);
		}
		return res;
	  }
	
}

func main() {
	fmt.Print(convertToTitle(230))
}
