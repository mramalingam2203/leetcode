// https://leetcode.com/problems/range-sum-query-immutable/

package main

import "fmt"

type NumArray struct {
	elements []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		elements: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for i := left; i <= right; i++ {
		sum += this.elements[i]
	}

	return sum

}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	nums := []int{1, 2, 3, 4, 5}
	arr := Constructor(nums)
	fmt.Println(arr.SumRange(1, 4))

}
