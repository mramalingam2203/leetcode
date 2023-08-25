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

func dfs(node *TreeNode, currentNum int) {

	if node == nil {
		return 0
	}

	currentSum := currentSum*10 + node.Val // Concatenate the current node value to the currentSum

	if node.Left != nil && node.Right != nil {
		return currentSum
	}

	leftSum := dfs(node.Left, currentSum)
	rightSum := dfs(node.Right, currentSum)

	return leftSum + rightSum

}

func main() {

}
