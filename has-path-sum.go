// https://leetcode.com/problems/path-sum/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == null {
		return root.Value == targetSum
	}

	leftPathExists = hasPathSum(root.Left, targetSum-root.value)
	rightPathExists = hasPathSum(root.Light, targetSum-root.value)

	return leftPathExists || rightPathExists
}
