package backtracking

/*
491. Increasing Subsequences
https://leetcode.com/problems/increasing-subsequences/
*/
func findSubsequencesRecursiveBacktracking(nums []int) [][]int {
	res := [][]int{}

	var hasSameElement = func(a [][]int, b []int) bool {
		for _, el := range a {
			foundDuplicate := true
			if len(el) != len(b) {
				foundDuplicate = false
				continue
			}
			for i := 0; i < len(el); i++ {
				if el[i] != b[i] {
					foundDuplicate = false
					break
				}
			}
			if foundDuplicate {
				return true
			}
		}

		return false
	}

	var backtracking func(startIndex int, current []int)
	backtracking = func(startIndex int, current []int) {
		if startIndex >= len(nums) {
			if len(current) > 1 && !hasSameElement(res, current) {
				cpy := make([]int, len(current))
				copy(cpy, current)
				res = append(res, cpy)
			}
			return
		}
		backtracking(startIndex+1, current)
		if len(current) == 0 || current[len(current)-1] <= nums[startIndex] {
			backtracking(startIndex+1, append(current, nums[startIndex]))
		}
	}

	backtracking(0, nil)
	return res
}

func findSubsequencesIterativeBacktracking(nums []int) [][]int {
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

func FindSubsequencesIterativeBacktracking(nums []int) [][]int {
	return findSubsequencesIterativeBacktracking(nums)
}

func FindSubsequencesRecursiveBacktracking(nums []int) [][]int {
	return findSubsequencesRecursiveBacktracking(nums)
}
