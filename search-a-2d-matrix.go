https://leetcode.com/problems/search-a-2d-matrix/

package main


func main() {

	array := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	searchMatrix(array)
}

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
    n := len(matrix[0])

    left := 0
    right := m * n - 1

	

    
}