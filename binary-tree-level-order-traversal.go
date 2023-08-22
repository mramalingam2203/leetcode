// https://leetcode.com/problems/binary-tree-level-order-traversal/

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

	levelOrder(root)

}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{{}}
	if root == nil {
		return nil
	}

	q := Queue{}
	q.Enqueue(*root)

	for !q.IsEmpty() {
		//current :=
		fmt.Println(q.Dequeue().Val)
		//fmt.Println(current.Val)

	}

	return result
}
