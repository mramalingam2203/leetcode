// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	val := 5
	_, _ = array, val
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	inorderIndex := findIndex(inorder, rootValue)
	root.Left = buildTree(preorder[1:inorderIndex+1], inorder[0:inorderIndex])
	root.Right = buildTree(preorder[inorderIndex+1:], inorder[inorderIndex+1:])

	return root

}

func findIndex(array []int, value int) int {
	var result int
	for i := 0; i < len(array); i++ {
		if array[i] == value {
			result = i
			break
		}
	}
	return result
}
