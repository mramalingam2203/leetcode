package main

import (
	"os"
	"fmt"
)

func specialPerm(nums []int) int {
	

	if len(nums) < 2 || len(nums) > 14{
		fmt.Println("Array length not good")
	}

	for num := range nums{
		if nums[num] < 1 || nums[num] > 109{
			fmt.Println("Numbers out of range"
		}
	}
    
}

func main(){

}