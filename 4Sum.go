package main

import (
	"fmt"
	_ "os"
)

func fourSum(nums []int, targetSum int) [][]int {
	quadruplets := [][]int{}
	numCount := make(map[int]int)

	// Count the occurrences of each number
	for _, num := range nums {
		numCount[num]++
	}

	// Iterate through the numbers and find the quadruplets
	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			// Skip duplicates
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k := j + 1; k < len(nums)-1; k++ {
				// Skip duplicates
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}

				remaining := targetSum - (nums[i] + nums[j] + nums[k])

				// Check if the remaining value is present in the count map
				if count, exists := numCount[remaining]; exists {
					// Ensure that we have enough occurrences of the remaining value
					if (remaining == nums[i] || remaining == nums[j] || remaining == nums[k]) && count < 2 {
						continue
					}

					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[k], remaining})
				}
			}
		}
	}

	
|

func main() {
	nums := []int{2, 2, 2, 2, 2}
	targetSum := 8

	quadruplets := findQuadruplets(nums, targetSum)

	fmt.Println("Quadruplets with a sum of", targetSum, ":")
	for _, quad := range quadruplets {
		fmt.Println(quad)
	}
}
