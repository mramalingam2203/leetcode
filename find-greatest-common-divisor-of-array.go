// https://leetcode.com/problems/find-greatest-common-divisor-of-array/

package main

import (
	"fmt"
	"sort"
)

func findGCD(nums []int) int {
	if len(nums) < 2 || len(nums) > 1000 {
		return 0
	}

	for _, i := range nums {
		if nums[i] < 1 || nums[i] > 1000 {
			return 0
		}
	}

	sort.Ints(nums)
	fmt.Println(nums)
	return gcd(nums[0], nums[len(nums)-1])

}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a

}

func main() {
	list := []int{3, 3}
	fmt.Println(findGCD(list))
}
