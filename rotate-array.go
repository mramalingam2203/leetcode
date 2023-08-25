// https://leetcode.com/problems/rotate-array/

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(array, 3)

}

func rotate(nums []int, k int) {

	fmt.Println(nums[0:k])

}
