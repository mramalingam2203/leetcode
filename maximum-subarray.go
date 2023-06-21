package main

import "fmt"

func findSubarrays(arr []int) int {
	temp := 0
	n := len(arr)
	sum := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum = 0
			subarray := arr[i : j+1]
			//subarrays = append(subarrays, subarray)
			for k := i; k < len(subarray); k++ {
				sum += subarray[k]
			}
			if sum > temp {
				temp = sum
			}
		}
	}

	return sum
}

func main() {
	// Example usage
	arr := []int{1, 2, 3}
	fmt.Println(findSubarrays(arr))

}
