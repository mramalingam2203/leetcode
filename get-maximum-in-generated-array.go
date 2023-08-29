// https://leetcode.com/problems/get-maximum-in-generated-array/

package main

import "fmt"

func getMaximumGenerated(n int) int {
	if n < 0 || n > 100 {
		return 0
	}

	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1

	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i/2] + nums[(i/2)+1]
		}
	}

	return maximum(nums)
}

func maximum(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	fmt.Println(getMaximumGenerated(7))
	// 0,1,1,2,1,3,2,3

}
