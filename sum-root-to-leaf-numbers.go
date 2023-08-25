// https://leetcode.com/problems/sum-root-to-leaf-numbers/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {

	return dfs(root, 0)

}

func dfs(node *TreeNode, currentSum int) int {

	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	leftSum := dfs(node.Left, currentSum)
	rightSum := dfs(node.Right, currentSum)

	return leftSum + rightSum

}

func main() {

	root := TreeNode{1, nil, nil}
	root.left = TreeNode{2}
	root.right = TreeNode{3}
	result = sumNumbers(root)
	print("Sum of root-to-leaf numbers:", result)

}
