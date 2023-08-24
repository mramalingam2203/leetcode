// https://leetcode.com/problems/palindrome-partitioning/


package main

func main() {

	fmt.Println(isPalindrome("abcddcba", 1, 4))

}

func isPalindrome(s, i, j){
    for i < j{
        if s[i] != s[j]{
            return false
		}
        i++, j++
	}
    return true

}
