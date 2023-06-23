package main

import "fmt"

func main() {
	// Example 9x9 array
	array9x9 := [9][9]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14, 15, 16, 17, 18},
		{19, 20, 21, 22, 23, 24, 25, 26, 27},
		{28, 29, 30, 31, 32, 33, 34, 35, 36},
		{37, 38, 39, 40, 41, 42, 43, 44, 45},
		{46, 47, 48, 49, 50, 51, 52, 53, 54},
		{55, 56, 57, 58, 59, 60, 61, 62, 63},
		{64, 65, 66, 67, 68, 69, 70, 71, 72},
		{73, 74, 75, 76, 77, 78, 79, 80, 81},
	}

	// Extract 3x3 subarrays
	subarrays := extractSubarrays(array9x9)

	// Print the subarrays
	for i, subarray := range subarrays {
		fmt.Printf("Subarray %d:\n", i+1)
		printSubarray(subarray)
		fmt.Println()
	}
}

// Function to extract all 3x3 subarrays from a 9x9 array
func extractSubarrays(array [9][9]int) [][3][3]int {
	subarrays := make([][3][3]int, 0)

	for i := 0; i <= 6; i += 3 {
		for j := 0; j <= 6; j += 3 {
			subarray := [3][3]int{}

			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					subarray[x][y] = array[i+x][j+y]
				}
			}

			subarrays = append(subarrays, subarray)
		}
	}

	return subarrays
}

// Function to print a 3x3 subarray
func printSubarray(subarray [3][3]int) {
	for _, row := range subarray {
		for _, val := range row {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}
}
