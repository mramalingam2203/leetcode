// https://leetcode.com/problems/number-of-employees-who-met-the-target/

package main

import "fmt"

func main() {

	hrs := []int{0, 1, 2, 3, 4}
	tgt := 2
	fmt.Println(numberOfEmployeesWhoMetTarget(hrs, tgt))

}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {

	if len(hours) < 1 || len(hours) > 50 || target > 1e5 {
		return 0
	}

	count := 0
	for i := range hours {
		if hours[i] < 0 || hours[i] > 1e5 {
			return 0
		}
		if hours[i] >= target {
			count++
		}
	}
	return count
}
