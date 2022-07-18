package backtracking

import (
	"sort"
	"fmt"
)

/*
698. Partition to K Equal Sum Subsets
https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
*/
func canPartitionKSubsets(nums []int, k int) bool {
	l := len(nums)
	if k > l {
		return false
	}
	s := 0
	for i := 0; i < l; i++ {
		s += nums[i]
	}
	if s%k != 0 {
		return false
	}
	sort.Ints(nums)
	expectedSum := s / k

	used := make([]int, len(nums))

	createKey := func(s []int) string {
		return fmt.Sprintf("%q", s)
	}

	memo := make(map[string]bool)
	
	var backtracking func(currentIndex int, currentSum int, piecesLeft int) bool

	backtracking = func(currentIndex int, currentSum int, piecesLeft int) bool {
		key := createKey(used)
		if piecesLeft == 0 {
			return true
		}
		if currentSum == expectedSum {
			res := backtracking(0, 0, piecesLeft-1)
			memo[key] = res
			return res
		}

		if val, ok := memo[key]; ok {
			return val
		}

		if currentSum > expectedSum || currentIndex >= l {
			return false
		}

		if used[currentIndex] == 0 {
			used[currentIndex] = 1
			if backtracking(currentIndex+1, currentSum+nums[currentIndex], piecesLeft) {
				return true
			}
			used[currentIndex] = 0
		}
		return backtracking(currentIndex+1, currentSum, piecesLeft)

	}
	return backtracking(0, 0, k-1)
}

func CanPartitionKSubsets(nums []int, k int) bool {
	return canPartitionKSubsets(nums, k)
}
