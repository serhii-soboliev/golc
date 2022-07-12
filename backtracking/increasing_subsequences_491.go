package backtracking

/*
491. Increasing Subsequences
https://leetcode.com/problems/increasing-subsequences/
*/
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var backtracking func(int)
	backtracking = func(startIndex int) {
		lc := len(cur)
		if lc > 1 {
			cpy := make([]int, lc)
			copy(cpy, cur)
			res = append(res, cpy)
		}
		used := map[int]bool{}
		for i := startIndex; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(cur) == 0 || cur[len(cur)-1] <= nums[i] {
				used[nums[i]] = true
				cur = append(cur, nums[i])
				backtracking(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtracking(0)
	return res
}

func FindSubsequences(nums []int) [][]int {
	return findSubsequences(nums)
}
