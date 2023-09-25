// https://leetcode.com/problems/h-index/

package main

import (
	"fmt"
	"sort"
)

func main() {
	citations := []int{1, 3, 1}
	fmt.Println(hIndex(citations))
}

func hIndex(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	hindex := 0

	for i, citationCount := range citations {
		// If the current citation count is greater than or equal to the current H-index, increment the H-index
		if citationCount >= i+1 {
			hindex = i + 1
		} else {
			// If the current citation count is less than the current H-index, stop iterating
			break
		}
	}
	return hindex
}
