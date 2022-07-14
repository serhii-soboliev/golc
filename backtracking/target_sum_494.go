package backtracking

/*
494. Target Sum
https://leetcode.com/problems/target-sum/
*/

func findTargetSumWays(nums []int, target int) int {
	count := 0
	var backtracking func(int, int)
	backtracking = func(startIndex int, sum int) {
		if startIndex == len(nums) {
			if sum == target {
				count += 1
			}
			return
		}
		backtracking(startIndex+1, sum-nums[startIndex])
		backtracking(startIndex+1, sum+nums[startIndex])

	}
	backtracking(0, 0)
	return count
}

func FindTargetSumWays(nums []int, target int) int {
	return findTargetSumWays(nums, target)
}
