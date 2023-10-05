// https://leetcode.com/problems/assign-cookies/

package main

import (
	"fmt"
	"sort"
)

func main() {

	g := []int{1, 2, 3}
	s := []int{3}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	childIndex := 0
	cookieIndex := 0
	satisfiedChildren := 0

	for childIndex < len(g) && cookieIndex < len(s) {
		// If the current cookie can satisfy the current child
		if s[cookieIndex] >= g[childIndex] {
			// Assign the cookie to the child
			satisfiedChildren++
			// Move to the next child and next cookie
			childIndex++
			cookieIndex++
		} else {
			cookieIndex++
		}
	}

	return satisfiedChildren
}
