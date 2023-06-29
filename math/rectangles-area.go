// https://leetcode.com/problems/rectangle-area/

package main

import (
	"fmt"
	_ "os"
)

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	var area1, area2, x_dist, y_dist, area_I int
	area1 = abs(ax1-ax2) * abs(ay1-ay2)
	area2 = abs(bx1-bx2) * abs(by1-by2)

	x_dist = min(ax2, bx2) - max(ax1, bx1)
	y_dist = min(ay2, ay1) - max(by2, by1)

	//area_I = 0
	if x_dist > 0 && y_dist > 0 {
		area_I = x_dist * y_dist
	}
	fmt.Println(x_dist, y_dist, area_I)

	return (area1 + area2 - area_I)
}

func abs(a int) {
	if a < 0 {
		return -1 * a
	}

	return a
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
}
