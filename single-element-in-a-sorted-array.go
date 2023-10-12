// https://leetcode.com/problems/single-element-in-a-sorted-array/

package main

import "fmt"

func main() {

	array := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(array))

}

func singleNonDuplicate(nums []int) int {

	freqMap := make(map[int]int)

	for num := range nums {
		freqMap[nums[num]]++
	}

	for key, value := range freqMap {
		if value == 1 {
			return key
		}
	}

	return -1
}
