// https://leetcode.com/problems/interleaving-string/

package main

func main() {

	str_1 = "aabcc"
	str_2 = "dbbca"
	str_3 = "aadbbcbcac"

	isInterleave(str_1, str_2, str_3)

}

func isInterleave(s1 string, s2 string, s3 string) bool {

	if len(s3) != len(s1)+len(s2) {
		return false
	}

}
