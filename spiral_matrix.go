package main

import "fmt"

func spiral_simulation(array [][]int) []int {
	m := len(array)    // number of rows
	n := len(array[0]) // number of columns

	ans := []int{}

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	x := 0
	y := 0
	di := 0

	// Create and initialize the 2D slice with zeros
	seen := make([][]bool, m)

	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var newX, newY int
	for i := 0; i < m*n; i++ {
		ans = append(ans, array[x][y])
		seen[x][y] = true

		newX = x + dr[di]
		newY = y + dc[di]

		if 0 <= newX && newX < m && 0 <= newY && newY < n && !seen[newX][newY] {
			x = newX
			y = newY
		} else {
			di = (di + 1) % 4
			x += dr[di]
			y += dc[di]
		}

		fmt.Println(x, y)
	}

	return ans

}

func main() {

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiral_simulation(matrix))
}
