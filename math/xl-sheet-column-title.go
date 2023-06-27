// https://leetcode.com/problems/excel-sheet-column-title


package main

fmt(
	"os"
	"fmt"
)



func convertToTitle(columnNumber int) string {
	var s string
	var i, rem int 

	for ;;;{
		rem = columnNumber%26
		if rem == 0 {
			s = s(append, "Z")
			columnNumber = (columnNumber/26)-1
		}else{
			s = s(append, strconv.Itoa(rem-1) + "A")
			columnNumber /= 26
		}
	}

	fmt.Print(s)
    
}


func main(){
	convertToTitle(23)
}