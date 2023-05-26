package main

import (
	"fmt"
	"sort"
)

func generateTriplets(nums []int) [][]int {
	for i:=0;i<len(nums);i++ {
		if nums[i] <= 


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
	result_1 := [][]int{}
	for i := 0; i < len(result); i++ {
		if result[i][0]+result[i][1]+result[i][2] == 0 {
			result_1 = append(result_1, result[i])
		}
	}

	return result_1
}

func main() {
	nums := []int{0, 1, 1}
	triplets := generateTriplets(nums)
	fmt.Println(triplets)
}
