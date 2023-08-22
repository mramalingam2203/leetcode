// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	preorder := []int{3, 9, 20, 15, 7}
	postorder := []int{9, 3, 15, 20, 7}
	fmt.Println(buildTree(preorder, postorder))
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(postorder) == 0 {
		return nil
	}

	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootValue}
	inorderIndex := findIndex(inorder, rootValue)

	leftSubtreeSize := inorderIndex

	root.Left = buildTree(inorder[0:leftSubtreeSize], postorder[0:leftSubtreeSize])
	root.Right = buildTree(inorder[leftSubtreeSize+1:], postorder[leftSubtreeSize:len(postorder)-1])

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
