// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import (
	"fmt"
	"os"
)

func main() {

	array := []int{2, 3, 4}
	number := 6
	fmt.Println(twoSum(array, number))

}

func twoSum(numbers []int, target int) []int {

	if len(numbers) < 2 || len(numbers) > 3e4 || target < -1000 || target > 1000 {
		fmt.Println("lenghts invalid")
		os.Exit(0)
	}
	indices := []int{}
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				if i != j {
					indices = append(indices, i+1)
					indices = append(indices, j+1)
					return indices
				}
			}
		}
	}
	return indices

}
