// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {

	freqMap := make(map[int]int)

	for num := range nums {
		freqMap[num]++
	}

	for key, val := range freqMap {
		if val > 1 {
			return true
		}
	}

	return false
}

func main() {{
	nums := []int[1,2,3,1]

	fmt.Println(containsDuplicate(nums))
}