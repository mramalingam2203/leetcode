// https://leetcode.com/problems/largest-divisible-subset/

package main

import "fmt"

func main() {

	array := []int{1, 2, 4, 8}
	fmt.Println(largestDivisibleSubset(array))

}

func largestDivisibleSubset(nums []int) []int {
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]%nums[j] == 0 || nums[j]%nums[i] == 0 {
				arr = append(arr, nums[i])
				arr = append(arr, nums[j])
				fmt.Println(nums[i], nums[j])
			}
		}
	}
	return arr
}
