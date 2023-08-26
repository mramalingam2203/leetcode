// https://leetcode.com/problems/rotate-array/

package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}

	reverse(array, 0, len(array)-1)
	//rotate(array, 1)
	fmt.Println(array)

}

/*

func rotate(nums []int, k int) {

	n := len(nums)

	if n < 1 || n > 1e5 || k < 0 || k > 1e5 {
		return
	}

	s1 := nums[n-k : n]
	s2 := nums[:k+1]

	for i := 0; i < len(s2); i++ {
		s1 = append(s1, s2[i])
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = s1[i]
	}
}
*/

func reverse(nums []int, start, end int) {

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}

}
