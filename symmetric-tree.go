// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)

}



function isMirror(left, right):
    if left == nil && right == null{
        return true
	}
    if left == null || right == null{
        return false
	}
    
    return (left.value = right.value) &&
           isMirror(left.Left, right.Right) &&
           isMirror(left.Right, right.Left)

}