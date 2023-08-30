//https://leetcode.com/problems/happy-number/

package main

import "strconv"

func main() {
	isHappy(19)
}

func isHappy(n int) bool {

	for n > 1 {
		s := strconv.Itoa(n)

		for i := 0; i < len(s); i++ {
			temp, _ := strconv.Atoi(s[i])
			sum := temp * temp
		}
		n = sum
	}

}
