// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import "fmt"

func main() {

	array := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(array))

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices)-1; i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
