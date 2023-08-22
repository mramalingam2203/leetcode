// https://leetcode.com/problems/binary-tree-preorder-traversal/

package main

import (
	"fmt"
)

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
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println("Preorder Traversal:")
	result := preorderTraversal(root)
	fmt.Println(result) // Output: [1 2 4 5 3 6 7]
}

func preorderTraversal(root *TreeNode) []int {

}
