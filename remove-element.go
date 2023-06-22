//https://leetcode.com/problems/remove-element/

package main

import (
	"fmt"
	"os"
)

func removeElement(nums []int, val int) int {
	if len(nums) < 0 || len(nums) > 100 {
		fmt.Println("lenght off range")
		os.Exit(0)
	}

	if val < 0 || val > 50 {
		fmt.Println("target off range")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] > 50 {
			fmt.Println("numbers off range")
			os.Exit(0)
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	if nums[len(nums)-1] == val {
		nums = myslice[:len(myslice) - 1]
	}

	return len(nums)
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	var removed int = 2
	fmt.Println(removeElement(arr, removed)) // remove

}
