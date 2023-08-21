// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

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

	fmt.Println(isSymmetric(root)) // Output: true

}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

}

func isMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)

}
