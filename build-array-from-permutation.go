package main

import (
	"fmt"
	"os"
)

func buildArray(nums []int) []int {

	/*
    1 <= nums.length <= 1000
    0 <= nums[i] < nums.length
    The elements in nums are distinct.
	*/
	if len(nums) < 1 || len(nums) > 1000{
		fmt.Println("length not a good one")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 ||  nums[i] > len(nums){
			fmt.Println("index out of range")
			os.Exit(0)
		}
	}

	


		
	}

	result := []int{}
	for i := range nums {
		result = append(result, nums[nums[i]])
	}

	return result

}

func main() {
	a := []int{5, 0, 1, 2, 3, 4}
	fmt.Println(buildArray(a))
}
