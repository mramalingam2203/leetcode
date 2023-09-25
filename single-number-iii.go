// https://leetcode.com/problems/single-number-iii/

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))

}

func singleNumber(nums []int) []int {

	frequency := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		frequency[nums[i]]++
	}

	result := []int{}

	// Iterate over all keys
	for key, val := range frequency {
		if val == 1 {
			result = append(result, key)
		}
	}

	return result
}
