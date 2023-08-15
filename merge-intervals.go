// https://leetcode.com/problems/merge-intervals/

package main

import "sort"

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) []Interval {

	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals based on their start times
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var merged []Interval
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		previous := merged[len(merged)-1]

		if current.Start <= previous.End {
			if current.End > previous.End {
				previous.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}



func merge(intervals [][]int) [][]int {

	intervalStruct := make([]Interval, len(intervals)
	for i := range intervals {
		intervalStruct[i].Start = intervals[i][0]
		intervalStruct[i].End = intervals[i][1]
	}

	mergeIntervals(intervalStruct []Interval) []Interval {

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
