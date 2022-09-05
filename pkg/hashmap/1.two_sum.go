package hashmap

/*
1. Two Sum
https://leetcode.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{i,j}
		}
		m[n] = i
	}
	panic("No solution")
}