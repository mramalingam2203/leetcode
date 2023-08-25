// https://leetcode.com/problems/gas-station/

package main

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	canCompleteCircuitCheck(gas, cost)
}

func canCompleteCircuitCheck(gas, cost []int) int {

	totalGas := 0
	totalCost := 0
	tank := 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]

		if tank < 0 {
			tank = 0
			startStation = i + 1
		}
	}

	if totalGas >= totalCost {
		return startStation
	} else {
		return -1
	}

}
