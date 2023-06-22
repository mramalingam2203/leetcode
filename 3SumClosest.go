package main

import "fmt"

func threeSumClosest(nums []int, target int) int {
	//var sum int
	results := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := []int{nums[i], nums[j], nums[k]}
				results = append(results, triplet)
			}
		}
	}
	temp := 0

	for i := 0; i < len(results); i++ {
		if results[i][0]+results[i][1]+results[i][2] > temp {
			temp = results[i][0] + results[i][1] + results[i][2]
		}
	}

	return temp

}

func mag(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	nums := []int{1, 1, 1, 0}
	fmt.Println(threeSumClosest(nums, -100))

}
