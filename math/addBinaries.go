package main

// fmt package provides the function to print anything
import (
	"fmt"
	"os"
)

// function which will add the binary strings
func addBinary(string1, string2 string) string {

	if len(string1) < 1 || len(string2) < 1 || len(string1) > 1e4 || len(string2) > 1e4 {
		fmt.Println("String length out of range")
		os.Exit(0)
	}
	/*
		for char := 0; char < len(string1); char++ {
			if string1[char] != 0 || string1[char] != 1 {
				fmt.Println("String a has strange character")
				os.Exit(0)
			}
		}

		for char := 0; char < len(string2); char++ {
			if string2[char] != 0 || string2[char] != 1 {
				fmt.Println("String b has strange character")
				os.Exit(0)
			}
		}
	*/
	// checking if the length of the first string is greater then
	// second then calling the function by swapping the parameters
	if len(string1) > len(string2) {
		return addBinary(string2, string1)
	}
	// finding the difference between the length of the strings
	difference := len(string2) - len(string1)

	// making both strings equal by adding 0 in front of a smaller string
	for i := 0; i < difference; i++ {
		string1 = "0" + string1
	}
	// initializing a variable carry to keep the track of carry after

	// each addition
	carry := "0"

	// In this variable we will store our final string
	answer := ""

	for i := len(string1) - 1; i >= 0; i-- {
		if string1[i] == '1' && string2[i] == '1' {
			if carry == "1" {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				carry = "1"
			}
		} else if string1[i] == '0' && string2[i] == '0' {
			if carry == "1" {
				answer = "1" + answer
				carry = "0"
			} else {
				answer = "0" + answer
			}
		} else if string1[i] != string2[i] {
			if carry == "1" {
				answer = "0" + answer
			} else {
				answer = "1" + answer
			}
		}
	}
	if carry == "1" {
		answer = "1" + answer
	}
	return answer
}
func main() {

	// declaring the strings
	var string1, string2 string

	// initializing the strings
	string1 = "10101"
	string2 = "100111"

	result := addBinary(string1, string2)
	// Printing the result of the addition of both the binary strings
	fmt.Println("The Numeric representation of", string1, "is", "21.")
	fmt.Println("The Numeric representation of", string2, "is", "39.")

	fmt.Println("The Binary addition of", string1, "and", string2, "is", result, "whose value in numeric is 60.")
}
