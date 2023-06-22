package main

import (
	"fmt"
	"os"
)

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2

	if n < 1 || n > 100 {
		fmt.Println("array index out of range")
		os.Exit(0)
	}

	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] > 100 {
			fmt.Println("Certani number out of range")
			os.Exit(0)
		}
	}

	// Find the first decreasing element from the right
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// Find the smallest element greater than nums[i] on its right side
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}

		// Swap nums[i] and nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Reverse the elements from i+1 to the end
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {
	nums := []int{1, 1, 5}

	// Find the next permutation
	nextPermutation(nums)

	fmt.Println(nums) // Output: [1 3 2]
}
