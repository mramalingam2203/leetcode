// Given a binary tree, determine if it is height-balanced

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	left  *TreeNode
	right *TreeNode
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
	fmt.Println(isBalanced(root)) // Output: 2
}

func isBalanced(root *TreeNode) bool {
	if root ==   nil{
        return true
	}
    
    leftHeight := getHeight(root.left)
    rightHeight := getHeight(root.right)
    
    if abs(leftHeight - rightHeight) > 1{
        return false
	}
    
    return isBalanced(root.left) AND isBalanced(root.right)
}



function getHeight(node *TreeNode){
    if node == nil {
        return 0
	}
    
    leftHeight := getHeight(node.left)
    rightHeight := getHeight(node.right)
    
    return 1 + max(leftHeight, rightHeight)

}

func abs( a int ) int {
	if a < 0{
		return -1*a
	}

	return a
}
