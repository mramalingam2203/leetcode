package main

import "fmt"

func threeSumClosest(nums []int, target int) int {

	results := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		var temp int
		sum := 0
		temp = sum
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := []int{nums[i], nums[j], nums[k]}
				results = append(results, triplet)
				sum = triplet[0] + triplet[1] + triplet[2]

			}
		}
	}

}

func mag(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	nums := []int{-1, 2, 1, -4}
	triplets := threeSumClosest(nums, 1)
	_ = triplets
	//for _, triplet := range triplets {
	//	fmt.Println(triplet)
	//}
}
