package backtracking

/*
494. Target Sum
https://leetcode.com/problems/target-sum/
*/
func findTargetSumWaysMemoBacktracking(nums []int, target int) int {
	l := len(nums)
	total := 0
	for _, n := range nums {
		total += n
	}

	memo := make([][]int, l)
	min := 0 - total - 1
	for i := range memo {
		memo[i] = make([]int, 2*total+1)
		for j := 0; j < 2*total+1; j++ {
			memo[i][j] = min
		}
	}

	var backtracking func(int, int) int

	backtracking = func(startIndex int, sum int) int {
		if startIndex == l {
			if sum == target {
				return 1
			} else {
				return 0
			}
		}
		if memo[startIndex][total+sum] != min {
			return memo[startIndex][total+sum]
		}
		add := backtracking(startIndex+1, sum+nums[startIndex])
		subtract := backtracking(startIndex+1, sum-nums[startIndex])
		memo[startIndex][total+sum] = add + subtract
		return memo[startIndex][total+sum]
	}

	return backtracking(0, 0)
}

func findTargetSumWaysBruteForce(nums []int, target int) int {
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

func FindTargetSumWaysMemoBacktracking(nums []int, target int) int {
	return findTargetSumWaysMemoBacktracking(nums, target)
}

func FindTargetSumWaysBruteForce(nums []int, target int) int {
	return findTargetSumWaysBruteForce(nums, target)
}
