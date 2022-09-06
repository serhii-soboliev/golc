package dynamicprogramming

/*
377. Combination Sum IV
https://leetcode.com/problems/combination-sum-iv/
*/
func combinationSum4(nums []int, target int) int {
    memo := make([]int, target + 1)
    for i := range memo {
        memo[i] = -1
    }

	var dp func(target int) int 

	dp = func(k int) int { 
		if k < 0 {
			return 0
		}
		if k == 0 {
			return 1
		}
		if memo[k] != -1 {
            return memo[k];
		}

		result := 0
		for _, v := range nums {
			result += dp(k - v)
		}
		memo[k] = result
		return result
	}

	return dp(target)
}

func CombinationSum4(nums []int, target int) int {
	return combinationSum4(nums, target)
}