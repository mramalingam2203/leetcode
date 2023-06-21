package main

import (
	"fmt"
	"os"
	"sort"
)

func threeSum(nums []int) [][]int {

	if len(nums) < 3 || len(nums) > 3000 {
		fmt.Println("array length out of range")
		os.Exit(0)
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] < -1e5 || nums[i] > 1e5 {
			fmt.Println("nnumbers no good")
			os.Exit(0)
		}
	}

	// Sort the array in ascending order
	sort.Ints(nums)

	var result [][]int

	// Iterate over each number as the potential first element of the triplet
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Use two pointers approach to find the other two elements
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Found a triplet that sums to zero
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate numbers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move the pointers inward
				left++
				right--
			} else if sum < 0 {
				// Sum is too small, move the left pointer to the right
				left++
			} else {
				// Sum is too large, move the right pointer to the left
				right--
			}
		}
	}

	return result
}

func main() {
	nums := []int{0, 0, 0}
	triplets := threeSum(nums)
	fmt.Println(triplets)
}
