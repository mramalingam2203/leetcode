// https: //leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

package main

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sortedArrayToBST(array)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildBST(nums, 0, len(nums)-1)

}

func buildBST(nums []int, start, end int) *TreeNode {

	if start > end {
		return nil
	}

	middleIndex := (start + end) / 2

	root := &TreeNode{
		Val: nums[middleIndex],
	}

	//root := TreeNode(nums[middleIndex])

	root.Left = buildBST(nums, start, middleIndex-1)
	root.Right = buildBST(nums, middleIndex+1, end)

	return root

}
