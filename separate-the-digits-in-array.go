package main

import (
	"strconv"
	"strings"
)

func separateDigits(nums []int) []int {

	digits := make([]string, len(nums))

	for i := 0; i < len(nums); i++ {
		digits[i] = strconv.Itoa(nums[i])
	}

	digits_1 := strings.Join(digits, "")
	results := strings.ReplaceAll(digits_1, " ", "")

	digits_int := make([]int, 0)

	for i := 0; i < len(results); i++ {
		k, _ := strconv.Atoi(string(results[i]))
		digits_int = append(digits_int, k)
	}
	return digits_int

}

func main() {

	list := []int{13, 25, 83, 77}

	separateDigits(list)
}
