// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import "fmt"

func main() {

	array := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(array))

}

func maxProfit(prices){
    if prices.length < 2:
        return 0
    
    maxProfit := 0
    
    for i := 1; i< len(prices)- 1; i++ {
        if prices[i] > prices[i - 1]{
            maxProfit += prices[i] - prices[i - 1]
		}
	}

    return maxProfit
}