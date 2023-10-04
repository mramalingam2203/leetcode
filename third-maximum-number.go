// https://leetcode.com/problems/third-maximum-number/

package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 3, 2, 2}
	fmt.Println(thirdMax(arr))

}

func thirdMax(nums []int) int {

	freq := make(map[int]int)
	array := []int{}

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
	}

	for key, _ := range freq {
		array = append(array, key)
	}

	// Sort the array in descending order
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})

	if len(array) <= 2 {
		return array[0]
	}
	return array[2]
}
