// https://leetcode.com/problems/132-pattern/

package main

import "fmt"

func main() {

	array := []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(array))

}

func find132pattern(nums []int) bool {

	if len(nums) < 3 {
		return false
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] < nums[j] {
				for k := j + 1; k < len(nums); k++ {
					if nums[i] < nums[k] && nums[k] < nums[j] {
						return true
					}
				}
			}
		}
	}

	return false
}
