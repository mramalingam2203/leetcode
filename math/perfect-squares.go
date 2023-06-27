// https://leetcode.com/problems/perfect-squares/

package main

import (
	"math"
)

func numSquares(n int) int {
	if math.Sqrt(n) - math.Floor(math.Sqrt(n)) == 0 {
		return 1}
	if n <= 3{
		return n
	}
	res := n
	var temp int
	for x:=1; x <= n; x++ {
		temp = x*x 
		if temp > n{
			break
		}else{
			res = min(res, 1 + numSquares(n-temp))
		}
		return res
	}
	

func main() {
	fmt.Println(numSquares(20))