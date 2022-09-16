package dynamicprogramming

/*
1770. Maximum Score from Performing Multiplication Operations
https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
*/

func maximumScoreDP1(nums []int, multipliers []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, m := len(nums), len(multipliers)
	memo := make([]map[int]int, n+1)
	for i := range memo {
		memo[i] = make(map[int]int)
	}

	var dp func(op, left int) int

	dp = func(op, left int) int {
		if op == m {
			return 0
		}
		if _, ok := memo[op][left]; ok {
			return memo[op][left]
		}

		l := nums[left]*multipliers[op] + dp(op+1, left+1)
		rigthIdx := n - 1 + left - op
		r := nums[rigthIdx]*multipliers[op] + dp(op+1, left)
		memo[left][op] = max(l, r)
		return memo[left][op]
	}
	return dp(0, 0)

}

func maximumScoreDP2(nums []int, multipliers []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n, m := len(nums), len(multipliers)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for op:=m-1; op >=0; op-- {
		for left := op; left>=0; left-- {
			l := nums[left]*multipliers[op] + dp[op+1][left+1]
			rigthIdx := n - 1 + left - op
			r := nums[rigthIdx]*multipliers[op] + dp[op+1][left]
			dp[op][left] = max(l,r)
		}
	}
	return dp[0][0]
}

func MaximumScoreDP2(nums []int, multipliers []int) int {
	return maximumScoreDP2(nums, multipliers)
}

func MaximumScoreDP1(nums []int, multipliers []int) int {
	return maximumScoreDP1(nums, multipliers)
}
