package main

import (
	"fmt"
	"strconv"
)

func separateDigits(nums []int) []int {

	digits := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		digits[i] = strconv.Itoa(nums[i])
	}

	/* 	digits_1 := strings.Join(digits, "")
	   	results := strings.ReplaceAll(digits_1, " ", "")
	*/
	//fmt.Println(digits_1)

	digits_int := make([]int, len(nums))

	for i := 0; i < len(digits_1); i++ {
		k, _ := strconv.Atoi(string(digits_1[i]))
		digits_int = append(digits_int, k)
	}

	fmt.Println(digits_int)

	return digits_int

}

func main() {

	list := []int{13, 25, 83, 77}

	separateDigits(list)
}
