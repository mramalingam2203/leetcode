// https: //leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, length(nums)-1)

}

func buildBST(nums, start, end) {

	if start > end {
		return nil
	}

	middleIndex := (start + end) / 2

	root = TreeNode(nums[middleIndex])

}

func main() {

}
