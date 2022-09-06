package contest

/*
2395. Find Subarrays With Equal Sum
https://leetcode.com/problems/find-subarrays-with-equal-sum/
*/
func FindSubarrays(nums []int) bool {
	n := len(nums)
	m := make(map[int]bool)

	for j := 0; j <= n-2; j++ {
		s := 0
		for k := j; k < j+2; k++ {
			s += nums[k]
		}
		if _, ok := m[s]; ok {
			return true
		}
		m[s] = true
	}

	return false
}

/*
2397. Maximum Rows Covered by Columns
https://leetcode.com/problems/maximum-rows-covered-by-columns/
*/
func MaximumRows(mat [][]int, cols int) int {
	m, n := len(mat), len(mat[0])
	rowSums := make(map[int]int)
	for i:=0; i<m; i++ {
		rowSums[i] = 0
		for j:=0; j<n; j++ {
			if mat[i][j] == 1 {
				rowSums[i] += 1
			}	
		}
	}
	alreadyCovered := 0
	for _,v := range rowSums {
		if v == 0 {
			alreadyCovered += 1
		}
	}

	result := alreadyCovered
	
	var backtracking func(pos int, colsUsedCount int, rowsCovered int) 
	
	backtracking = func(pos int, colsUsedCount int, rowsCovered int) { 
		if pos == n || colsUsedCount == cols {
			if rowsCovered > result {
				result = rowsCovered
			}
			return
		}
		for i:=pos; i<n; i++ {
			rowsDelta := 0

			rowSumsChanges := make([]int, m)
			for j:=0; j<m; j++ {
				if mat[j][i] == 1 && rowSums[j] > 0 {
					rowSums[j] -= 1
					rowSumsChanges[j] += 1
					if rowSums[j] == 0 {
						rowsDelta += 1
					}
				}
			}

			backtracking(i+1, colsUsedCount+1, rowsCovered + rowsDelta)

			for j:=0; j<m; j++ {
				rowSums[j] += rowSumsChanges[j]
			}
		}
	}

	backtracking(0, 0, alreadyCovered)

	return result
}

/*
2396. Strictly Palindromic Number
https://leetcode.com/problems/strictly-palindromic-number/
*/
func isStrictlyPalindromic(n int) bool {

	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	helper := func(n int, b int) bool {
		ans := ""
		for n>=b {
			temp := n%b
			n = n/b
			ans += string(rune(temp))
		}
		ans += string(rune(n))
		return ans == reverse(ans)

	}
	for i:=2; i<=n-2; i++ {
		if !helper(n, i){
			return false
		}
	}
	return true
}








