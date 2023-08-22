// https://leetcode.com/problems/binary-tree-level-order-traversal/

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
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	fmt.Println(levelOrder(root))

}

func levelOrder(root *TreeNode) [][]int {

	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]
			levelValues = append(levelValues, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		result = append(result, levelValues)
	}

	return result

}
