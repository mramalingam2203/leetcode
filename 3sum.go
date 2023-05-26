package main

import (
	"fmt"
	"os"
	"sort"
)

func generateTriplets(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		if nums[i] <= -100000 || nums[i] >= 100000 {
			fmt.Println("out of range")
			os.Exit(0)
		}
	}

	if len(nums) < 3 || len(nums) > 3000 {
		fmt.Println("lenght of array not good")
		os.Exit(0)
	}

	result := [][]int{}
	n := len(nums)

	// Helper function to generate combinations recursively
	var generateCombos func(index int, combo []int)
	generateCombos = func(index int, combo []int) {
		if len(combo) == 3 {
			// Found a triplet combination
			// Append a copy of the combo to the result
			temp := make([]int, 3)
			copy(temp, combo)
			result = append(result, temp)
			return
		}

		for i := index; i < n; i++ {
			// Skip duplicates
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			combo = append(combo, nums[i])
			generateCombos(i+1, combo)
			combo = combo[:len(combo)-1]
		}
	}

	// Sort the input array to skip duplicates efficiently
	sort.Ints(nums)
	generateCombos(0, []int{})
	/*
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
	*/
	result_1 := [][]int{}
	for i := 0; i < len(result); i++ {
		if result[i][0]+result[i][1]+result[i][2] == 0 {
			result_1 = append(result_1, result[i])
		}
	}

	return result_1
}

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	triplets := generateTriplets(nums)
	fmt.Println(triplets)
}
