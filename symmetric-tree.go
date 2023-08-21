// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	array := []int{1, 2, 2, 3, 4, 4, 3}
	isSymmetric(array[0])

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
