// https://leetcode.com/problems/decode-xored-array/

package main

import (
	"fmt"
)

func decode(encoded []int, first int) []int {
	key := encoded[0] ^ first

	for i := 0; i < len(encoded)-1; i++ {
		fmt.Println(key ^ encoded[i+1])
	}
	return encoded
}

func main() {
	ncoded := []int{1, 2, 3}
	start := 1

	fmt.Println(decode(ncoded, start))
}
