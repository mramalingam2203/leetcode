// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

package main

type TreeNode struct {
	Val   int
	Next *TreeNode
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left.Next = root.Right

	if root.left != nil {
		root.Left.Next = root.right
	}

	if root.Right != nil && root.Next != nil {

		root.Right.Next = root.Next.Left
	}

	connect(root.Left)
	connect(root.Right)
}
