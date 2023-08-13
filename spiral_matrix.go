package main

func spiral_simulation(array [][]int) {
	m := len(array)    // number of rows
	n := len(array[0]) // number of columns

	ans := make([]int, m*n)

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	x := 0
	y := 0
	di := 0

	for i := 0; i < m*n; i++ {
		ans = append(ans, array[i][j])

	}

}

func main() {

	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	spiral_simulation(matrix)
}
