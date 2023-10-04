// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(findDuplicates(nums))
}

func findDuplicates(nums []int) []int {

	result := []int{}
	freq := make(map[int]int)

	for i := range nums {
		freq[nums[i]]++
	}

	// Iterate through the map using a for loop and range
	for key, value := range freq {
		if value == 2 {
			result = append(result, key)
		}
	}

	return result
}
