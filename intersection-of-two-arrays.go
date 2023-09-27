// https://leetcode.com/problems/intersection-of-two-arrays/

package main

import "fmt"

func main() {

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{0, 2, 3, 5, 6}
	intersection(arr1, arr2)

}

func intersection(nums1, nums2 []int) []int {

	map1, map2 := make(map[int]int), make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		map1[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		map2[nums2[i]]++
	}
	// Create a slice to store common keys.
	commonKeys := []int{}

	// Iterate through the keys in map1.
	for key := range map1 {
		// Check if the key also exists in map2.
		if _, exists := map2[key]; exists {
			commonKeys = append(commonKeys, key)
		}
	}

	fmt.Println(commonKeys)
}
