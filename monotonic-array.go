// https://leetcode.com/problems/monotonic-array/?envType=study-plan-v2&envId=programming-skills

package main

func main() {
	nums := []int{1, 2, 3, 4}
	isMonotonic(nums)
}

func isMonotonic(nums []int) bool {
	temp := nums[1] - nums[0]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != temp {
			return false
		}
		temp = nums[i+1] - nums[i]
	}

	return true
}
