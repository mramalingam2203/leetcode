// https://leetcode.com/problems/find-the-duplicate-number/

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}

	fmt.Println(findDuplicate(array))

}

func findDuplicate(nums []int) int {

	n := len(nums)

	//if n < 1 || n > 1e5

	frequency := make(map[int]int)

	for i := 0; i < n; i++ {
		frequency[nums[i]]++
	}

	targetValue := 2
	var result int
	for key, value := range frequency {
		if value == targetValue {
			result = key
		}
	}

	return result
}
