// https://leetcode.com/problems/max-points-on-a-line/

package main

func main() {
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	maxPoints(points)
}

func maxPoints(points [][]int) int {

	if len(points) < 1 || len(points) > 300 {
		return 0
	}

	var slope, maxPoints int = 0, 0
	_ = maxPoints

	for i := 0; i < len(points); i++ {

		slopeCount := make(map[int]int)

		verticalPoints := 0

		for j := 0; j < len(points); j++ {

			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 {
				verticalPoints++
			} else {
				slope = dy / dx
				slopeCount[slope]++

			}

			maxPoints = max(maxPoints, duplicatePoints+verticalPoints)

		}

	}

	return 0

}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
