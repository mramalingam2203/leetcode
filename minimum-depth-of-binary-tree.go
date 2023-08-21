// Given the root of a binary tree, return its maximum depth.

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println(minDepth(root)) // Output: 2
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	return 1 + min(leftDepth, rightDepth)

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
