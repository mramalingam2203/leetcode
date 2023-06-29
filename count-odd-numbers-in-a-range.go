package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(countOdds(8, 10))
}

func countOdds(low int, high int) int {

	if low < 0 || low > 1e9 || high < 0 || high > 1e9 || low > high {
		fmt.Println("solution not defined")
		os.Exit(0)
	}
	count := 0
	for i := low; i < high; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count

}
