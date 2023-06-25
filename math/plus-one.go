// https://leetcode.com/problems/plus-one/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func convertIntArrayToInt(elems []int) int {
	b := ""
	for _, v := range elems {
		b += strconv.Itoa(v)
		b += ""
	}
	result, _ := strconv.Atoi(b)
	return result
}

func NumToArray(num int) []int {
	arr := make([]int, len(strconv.Itoa(num)))

	for i := len(arr) - 1; num > 0; i-- {
		arr[i] = num % 10
		num = int(num / 10)
	}
	fmt.Println(arr)
	return arr
}

func plusOne(digits []int) []int {
	if len(digits) == 0 || len(digits) > 100 {
		fmt.Println("int array lenngth out of range")
		os.Exit(0)
	}
	/*
		for i := 0; i < len(digits); i++ {
			if digits[i] < 1 || digits[i] > 9 {
				fmt.Println("numbers out of range")
				os.Exit(0)
			}
		}
	*/
	if digits[0] == '0' {
		fmt.Println("number invalid")
		os.Exit(0)
	}
	var plusone uint64
	number := convertIntArrayToInt(digits)
	plusone = uint64(number) + 1
	fmt.Println(plusone)
	//fmt.Println(uint64(number))
	return NumToArray(plusone)
}

func main() {
	a := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	plusOne(a)
}
