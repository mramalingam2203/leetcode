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

}

func main() {

}
