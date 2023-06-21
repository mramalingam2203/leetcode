package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
	var sum int
	results := [][]int{}
	var temp int = target + 1

	for i := 0; i < len(nums)-2; i++ {
		sum = 0
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := []int{nums[i], nums[j], nums[k]}
				results = append(results, triplet)
				sum = triplet[0] + triplet[1] + triplet[2]
				if mag(sum-target) < temp {
					temp = sum
				}
			}
		}
	}
	return sum
}

func mag(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))

}
