// https://leetcode.com/problems/kth-largest-element-in-an-array/

package main

import "container/heap"

type MinHeap []int

// MinHeap represents a min heap of integers.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	// Push an element onto the min heap
	heap.Push(h, 0)

	// Pop the minimum element from the min heap
	min := heap.Pop(h).(int)
	fmt.Prinstln("Minimum element:", min)

	// Peek at the minimum element without popping it
	min = (*h)[0]
	fmt.Println("Minimum element (peek):", min)
}
