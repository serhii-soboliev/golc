package backtracking

/*
1799. Maximize Score After N Operations
https://leetcode.com/problems/maximize-score-after-n-operations/
*/
func gcd(a, b int) int {
	if a < b {
		a,b = b,a
	}
	if a == 0 { return 0 }
	if b == 0 { return a }
	return gcd(a - b, b)
}

func gcds(nums []int) [][]int {
	n := len(nums)
	b := make([][]int, n)
	for i:=0; i<n; i++ {
		b[i] = make([]int, n)
	}
	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			curGcd := gcd(nums[i], nums[j])
			b[i][j] = curGcd
			b[j][i] = curGcd
		}
	}

	return b
}

func maxScore(nums []int) int {
	n := len(nums)
	memo := make(map[int]int)
	b := gcds(nums)

	var dfs func(pos int, mask int) int

	dfs = func(pos int, mask int) int {
		if pos > n / 2 {
			return 0
		}
		if v, ok := memo[mask];ok {
			return v
		}
		res := 0
		for i:=0; i<n; i++ {
			for j:=i+1; j<n; j++ {
				newMask := (1 << i) | (1 << j)
				if mask & newMask == 0 {
					sum := pos * b[i][j] + dfs(pos+1, newMask | mask)
					if sum > res {
						res = sum
					}
				}
			}	
		}
		memo[mask] = res
		return res
	}
	
	return dfs(1,0)
}

func MaxScore(nums []int) int {
	return maxScore(nums)	
}