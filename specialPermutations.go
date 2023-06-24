package main

import (
	"fmt"
	"os"
)

func specialPerm(nums []int) int {

	if len(nums) < 2 || len(nums) > 14 {
		fmt.Println("Array length not good")
		os.Exit(0)
	}

	for num := range nums {
		if nums[num] < 1 || nums[num] > 109 {
			fmt.Println("Numbers out of range")
			os.Exit(0)
		}
	}

	return 0
}

func main() {

}
