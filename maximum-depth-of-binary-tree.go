// Given the root of a binary tree, return its maximum depth.

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	// Example usage
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}

	fmt.Println(maxDepth(root))

}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)

}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
