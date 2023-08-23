// https://leetcode.com/problems/sort-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	 // Step 1: Divide the list
	 middle := findMiddle(head)
	 leftHalf := head
	 rightHalf := middle.next
	 middle.next := null

}

func findMiddle(head *ListNode) *ListNode {

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {

	node1 := &ListNode{Val: 1, Next: nil}
	node2 := &ListNode{Val: 2, Next: nil}
	node3 := &ListNode{Val: 3, Next: nil}
	node4 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 5, Next: nil}
	node6 := &ListNode{Val: 6, Next: nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	fmt.Println(findMiddle(node1))

}
