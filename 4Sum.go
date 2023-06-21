package main

import (
	"fmt"
	_ "os"
	"sort"
)

/*
func fourSum(nums []int, target int) [][]int {

}
*/

func findQuadruplets(nums []int, target int) [][]int {
	// Sort the array in ascending order
	sort.Ints(nums)

	quadruplets := [][]int{}
	length := len(nums)

	// Fix the first element
	for i := 0; i < length-3; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Fix the second element
		for j := i + 1; j < length-2; j++ {
			// Skip duplicates
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// Fix the third and fourth elements using two pointers
			left := j + 1
			right := length - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]

				if sum == 0 {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})

					// Skip duplicates
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				} else if sum < 0 {
					left++
				} else {
					right--
				}
			}
		}
	}

	return quadruplets
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2} // Example input array
	target := 3
	quadruplets := findQuadruplets(nums, target)

	fmt.Println("Quadruplets:")
	for _, quad := range quadruplets {
		fmt.Println(quad)
	}
}
