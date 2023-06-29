// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/?envType=study-plan-v2&envId=programming-skills

package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(a))
}

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	temp := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i+1]-arr[i]) != temp {
			return false
		}
		temp = arr[i+1] - arr[i]
		fmt.Println(temp)
	}
	return true

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
