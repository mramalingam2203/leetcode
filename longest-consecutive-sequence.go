// https://leetcode.com/problems/longest-consecutive-sequence/

package main

import(
	"fmt"
)

func main() {

//	array := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	array := []int{100,4,200,1,3,2}
	fmt.Println(longestConsecutive(array))

}


func longestConsecutive(nums []int) int {

	numSet := NewSet()

	// Create a set from the array
	for _, num := range nums {
		numSet.Add(num)
	}
	
	longestStreak := 0
	
	for _, num := range nums {
		if !numSet.Contains(num - 1) {
			currentNum := num
			currentStreak := 1

			for numSet.Contains(currentNum + 1) {

				currentNum++
				currentStreak++

			}
		
			if currentStreak > longestStreak{
					longestStreak = currentStreak
			}
	}
	
}				

return longestStreak

}



type Set map[int]bool

func NewSet() Set {
	return make(Set)
}

func (s Set) Add(item int) {
	s[item] = true
}

func (s Set) Contains(item int) bool {
	return s[item]
}

func (s Set) Remove(item int) {
	delete(s, item)
}

func (s Set) Size() int {
	return len(s)
}


