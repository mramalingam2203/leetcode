// https://leetcode.com/problems/find-greatest-common-divisor-of-array/

package main

import (
	"fmt"
	"sort"
)

func findGCD(nums []int) int {

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
	list := []int{7, 5, 6, 8, 3}
	fmt.Println(findGCD(list))
}
