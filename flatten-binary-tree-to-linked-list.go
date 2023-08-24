// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func flatten(root *TreeNode) {

	if root == nil {
		return
	}

	left := root.Left
	right := root.Right

	flatten(left)
	flatten(right)

	// Move the flattened left subtree to the right
	root.Left = nil
	root.Right = left

	end := root

	for end.Right != nil {

		end = end.Right
	}

	end.Right = right

}

}
