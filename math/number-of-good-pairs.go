// https://leetcode.com/problems/number-of-good-pairs/

package main

import (
	"fmt"
	"os"
)


void CombinationRepetitionUtil(int chosen[], int arr[],
                    int index, int r, int start, int end)
{
    // Since index has become r, current combination is
    // ready to be printed, print
    if (index == r)
    {
        for (int i = 0; i < r; i++)
            printf("%d ", arr[chosen[i]]);
        printf("\n");
        return;
    }
 
    // One by one choose all elements (without considering
    // the fact whether element is already chosen or not)
    // and recur
    for (int i = start; i <= end; i++)
    {
        chosen[index] = i;
        CombinationRepetitionUtil(chosen, arr, index + 1,
                                               r, i, end);
    }
    return;
}





func numIdenticalPairs(nums []int) int {
	if len(nums) < 1 || len(nums) > 100 {
		fmt.Printf("Array length out of range")
		os.Exit(0)
	}
	for i := 0; i < len(nums); i++ 
	// Utility function to generate combinations recursively
	func combinationUtil(arr []int, n, r, index int, data []int, dataIndex int) {
		if dataIndex == r {
			// Print the combination
			fmt.Println(data)
			return
		}
	
		if index >= n {
			return
		}
	
		data[dataIndex] = arr[index]
		combinationUtil(arr, n, r, index+1, data, dataIndex+1)
		combinationUtil(arr, n, r, index+1, data, dataIndex)
	}
	{
		if nums[i] < 1 || nums[i] > 100 {
			fmt.Println("number out of range")
			os.Exit(0)
		}

		
	}

	return 0
}

func main() {
	array := []int{1, 2, 3, 4}
	numIdenticalPairs(array)
}
