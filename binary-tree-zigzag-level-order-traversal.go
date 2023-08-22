//https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

package main

import (
	"fmt"
	"os"
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

	fmt.Println(levelOrder(root))

}

func levelOrder(root *TreeNode) [][]int {

	var result [][]int

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	var nTreeNodes int = 0
	for len(queue) > 0 {

		levelSize := len(queue)
		levelValues := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			if current.Val < -1000 || current.Val > 1000 {
				fmt.Println("Invalid node value")
				os.Exit(0)
			}

			levelValues = append(levelValues, current.Val)
			
			if i % 2 != 0
				if current.Left != nil {
					queue = append(queue, current.Left)
					nTreeNodes++
				}
				if current.Right != nil {
					queue = append(queue, current.Right)
					nTreeNodes++
				}
			}


			if i % 2 == 0
				if current.Right != nil {
					queue = append(queue, current.Right)
					nTreeNodes++
				}
				if current.Left != nil {
					queue = append(queue, current.Left)
					nTreeNodes++
				}
			}

			if nTreeNodes > 2000 {
				fmt.Println("tree too long")
				os.Exit(0)
			}
		}

		result = append(result, levelValues)
	}

	return result

}
