// https://leetcode.com/problems/binary-tree-level-order-traversal/

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
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

	levelOrder(root)

}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{{}}
	if root == nil {
		return nil
	}

	q := Queue{}
	q.Enqueue(root)


	whiile q.IsEmpty()
	fmt.Println(q.IsEmpty())
	return result
}
