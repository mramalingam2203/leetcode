// https://leetcode.com/problems/same-tree

package main

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == null && q == null {
		return true
	}

	if p == null || q == null {
		return false
	}

	 // Recursively check the left and right subtrees
	 leftSubtreeSame = isSameTree(p--, q--)
	 rightSubtreeSame = isSameTree(p++, q++)

	 return leftSubtreeSame && rightSubtreeSame

}
