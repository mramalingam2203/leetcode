// https://leetcode.com/problems/merge-intervals/

package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	// Use the sort.Slice function with a custom comparator
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := [][]int{{}}

	for i := 0; i < len(intervals)-1; i++ {

		if intervals[i][1] >= intervals[i+1][0] {

		} else {

			int1 := min(intervals[i][0], intervals[i+1][0])
			int2 := max(intervals[i][1], intervals[i+1][1])

			fmt.Println(int1, int2)

		}

	}

	return results

}

func min(a, b int) int {

	if a < b {
		return a
	}

	return b

}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b

}

func main() {

	list := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(list)

}
