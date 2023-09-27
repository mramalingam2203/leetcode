// https://leetcode.com/problems/top-k-frequent-elements/

package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 4}
	k := 3
	fmt.Println(topKFrequent(array, k))

}

func topKFrequent(nums []int, k int) []int {

	freqMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}

	// Create a reverse map where keys are frequencies and values are slices of keys.
	reverseMap := make(map[int][]int)
	for key, frequency := range freqMap {
		reverseMap[frequency] = append(reverseMap[frequency], key)
	}

	// Extract the top k frequent values.
	var topKFrequent []int
	for i := len(nums); i >= 0 && len(topKFrequent) < k; i-- {
		if keys, ok := reverseMap[i]; ok {
			sort.Ints(keys) // Sort keys with the same frequency.
			topKFrequent = append(topKFrequent, keys...)
		}
	}

	return topKFrequent[:k]
}
