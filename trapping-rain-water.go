// https://leetcode.com/problems/trapping-rain-water/

package main

import "fmt"

func main() {

	bars := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(bars))
}

func trap(height []int) int {

	n := len(height)

	if n < 1 || n > 2e4 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i <= n-1; i++ {
		if height[i] < 0 || height[i] > 1e5 {
			return 0
		}
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var trappedWater int = 0

	for i := 0; i < n; i++ {
		trappedWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return trappedWater

}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b

}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b

}
