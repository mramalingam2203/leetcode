// https://leetcode.com/problems/evaluate-reverse-polish-notation/

package main

import (
	"fmt"
	"os"
	"strconv"
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

func findOperand(char rune) bool {
	if char >= 48 && char <= 57 {
		return true
	} else {
		return false
	}
}

func findOperandI(s string) bool {
	atoi, _ := strconv.Atoi(s)
	if atoi < -200 && atoi > 200 {
		fmt.Println("Out of range")
		os.Exit(0)
	}
	return false

}

func main() {
	//tokens = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	runes := []rune("0123456789")
	fmt.Println(runes)

	for i := 0; i < len(runes); i++ {
		fmt.Println(findOperand(runes[i]))
	}

}
