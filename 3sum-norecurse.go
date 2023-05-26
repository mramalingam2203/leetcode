package main

import "fmt"

func generateTriplets(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	// Generate triplet combinations without repetition
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// Check for uniqueness
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					triplet := []int{nums[i], nums[j], nums[k]}
					result = append(result, triplet)
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	triplets := generateTriplets(nums)
	fmt.Println(triplets)
}
