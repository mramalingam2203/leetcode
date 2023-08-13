/*
Jump Game
Merge Intervals
Insert Interval
Spiral Matrix II
Unique Paths II
Minimum Path Sum
Set Matrix Zeroes
Search a 2D Matrix
Sort Colors
*/

package main

import (
	"fmt"
	"os"
)

func spiral_simulation(n int) []int {

	if n < 1 || n > 20 {
		fmt.Println("Invalid matrix size")
		os.Exit(0)
	}

	ans := []int{}

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	x := 0
	y := 0
	di := 0

	// Create and initialize the 2D slice with zeros
	seen := make([][]bool, n)

	for i := range seen {
		seen[i] = make([]bool, n)
	}

	var newX, newY int
	for i := 0; i < n*n; i++ {

		ans = append(ans, i)
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
	fmt.Println(spiral_simulation(3))
}
