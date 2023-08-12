// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

package main

import (
	"fmt"
	"os"
)

func main() {

	candys := []int{2, 8, 7}
	xtracandys := 1

	fmt.Println(kidsWithCandies(candys, xtracandys))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) < 2 || len(candies) > 100 || extraCandies < 1 || extraCandies > 50 {
		fmt.Println("invalid input")
		os.Exit(0)
	}

	results := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		now_in_hand := candies

		if candies[i] < 1 || candies[i] > 100 {
			fmt.Println("invalid inputs")
			os.Exit(0)
		}
		now_in_hand[i] += extraCandies
		if greatestIdx(now_in_hand) == i {
			results[i] = true
		}
		now_in_hand[i] -= extraCandies
	}
	return results

}

func greatestIdx(nums []int) int {
	max := nums[0]
	var max_idx int

	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
			max_idx = i
		}
	}

	return max_idx
}
