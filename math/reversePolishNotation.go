// https://leetcode.com/problems/evaluate-reverse-polish-notation/

package main

import (
	"fmt"
	_ "os"
)

/*
func evalRPN(tokens []string) int {

}

func findRPN(char int) bool{
	if
}

int findOperator(char ch){
	if(ch == '+'|| ch == '-'|| ch == '*'|| ch == '/' || ch == '^')
	   return 1;//if they found an operatAor
	return -1;//if they didn't found an operator
 }

*/

func findOperator(char rune) bool {
	if char == 42 || char == 43 || char == 45 || char == 47 {
		return true
	} else {
		return false
	}
}

func main() {
	//tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	runes := []rune("*+-#")
	for i := 0; i < len(runes); i++ {
		fmt.Println(findOperator(runes[i]))
	}
}
