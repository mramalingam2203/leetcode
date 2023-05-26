package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func main() {
	// Example usage
	arr := []int{1, 1, 2, 3, 3, 3, 4, 4, 5, 6, 6, 6}
	fmt.Println("Original array:", arr)

	newLength := removeDuplicates(arr)
	fmt.Println("New length:", newLength)
	fmt.Println("Updated array:", arr[:newLength])
}
