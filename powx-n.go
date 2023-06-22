package main

import (
	"fmt"
	"os"
)

func myPow(x float64, n int) float64 {
	constraints(x, n)
	var power float64 = x

	for i := 1; i < abs(n); i++ {
		power *= x
	}
	/*
		if n > 0 {
			if power < -1e4 || power > 1e4 {
				fmt.Print("power out of range")
				os.Exit(0)
			}
		} else if n < 0 {
			if power < -1e4 || power > 1e4 {
				os.Exit(0)
			}

		}
	*/
	switch {

	case n < 0:
		power = 1 / power

	case n == 0:
		power = 1.0

	}

	return power

}

func constraints(x float64, n int) {
	var two_power_thirtyone int = 2147483648

	if x < -100.0 || x > 100.0 {
		fmt.Print("x out of range")
		os.Exit(0)
	}

	if -1*two_power_thirtyone > n || two_power_thirtyone-1 < n {
		fmt.Print("condition 1 not met")
		os.Exit(0)
	}

}
func abs(N int) int {
	if N < 0 {
		N *= -1
	}
	return N
}

func main() {
	fmt.Println(myPow(34.00515, -3))
}
