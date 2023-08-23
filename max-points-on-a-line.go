// https://leetcode.com/problems/max-points-on-a-line/

func main() {
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
}

func maxPoints(points [][]int) int {

	if len(points) < 1 || len(points) > 300 {
		return 0
	}

	maxPoints := 0

	for i := 0; i < len(points); i++ {

		slopeCount := make(map[int]int)

		vertcalPoints := 0

		for j := 0; j < len(points); j++ {

			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 {
				vertcalPoints++
			} else {
				slope := dy / dx
			}

		}

		fmt.Println(slope)
	}

}