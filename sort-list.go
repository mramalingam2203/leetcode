// https://leetcode.com/problems/sort-list/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Step 1: Divide the list
	middle := findMiddle(head)
	leftHalf := head
	rightHalf := middle.Next
	middle.Next = nil

	// Step 2: Recursively sort the halves
	sortedLeft := sortList(leftHalf)
	sortedRight := sortList(rightHalf)

	// Step 3: Merge the sorted halves
	sortedList := mergeLists(sortedLeft, sortedRight)
	return sortedList

	return head

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

func mergeLists(head1, head2 *ListNode) *ListNode {

	var dummyHead *ListNode // Create a dummy head node
	current := dummyHead

	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			current.Next = head2
			head1 = head1.Next
		} else {
			current.Next = head2
			head2 = head2.Next
		}

		current = current.Next

	}

	// Attach remaining nodes if any
	if head1 != nil {
		current.next = head1
	} else {
		current.next = head2
	}

	return dummyHead.Next
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
