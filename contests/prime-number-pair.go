// https://leetcode.com/contest/weekly-contest-352/problems/prime-pairs-with-target-sum/

package main

import (
	"fmt"
	"os"
)

func main() {
	a := 10
	findPrimePairs(a)

}

func findPrimePairs(n int) [][]int {
	if n < 1 || n < 1e6 {
		fmt.Println("Numbers out of range")
		os.Exit(0)
	}
	prime_numbers := []int{}

	for i := 1; i < n; i++ {
		if sieveOfEratosthenes(i) == true {
			prime_numbers = append(prime_numbers, i)
		}
	}

	fmt.Println(prime_numbers)
}

func sieveOfEratosthenes(n int) []bool {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	// Loop through all numbers from 2 to square root of n
	for p := 2; p*p <= n; p++ {
		// If prime[p] is not changed, then it is a prime
		if prime[p] == true {
			// Update all multiples of p
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	return prime
}

func combinations(arr []int, k int) [][]int {
	if k <= 0 || len(arr) < k {
		return [][]int{}
	}

	combs := [][]int{}
	combination := make([]int, k)

	var generate func(start, pos int)
	generate = func(start, pos int) {
		if pos == k {
			comb := make([]int, k)
			copy(comb, combination)
			combs = append(combs, comb)
			return
		}

		for i := start; i < len(arr); i++ {
			combination[pos] = arr[i]
			generate(i+1, pos+1)
		}
	}

	generate(0, 0)
	return combs
}
