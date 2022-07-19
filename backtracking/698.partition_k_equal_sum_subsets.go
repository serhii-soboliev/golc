package backtracking

/*
698. Partition to K Equal Sum Subsets
https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
*/
import "sort"

func backtracking(startIndex int, currentSum int, piecesLeft int, expectedSum int, used *[]bool, nums *[]int) bool {
	l := len(*nums)
	if piecesLeft == 0 {
		return true
	}
	if currentSum == expectedSum {
		return backtracking(0, 0, piecesLeft-1, expectedSum, used, nums)
	}

	if currentSum > expectedSum || startIndex >= l {
		return false
	}

	for i := startIndex; i < l; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			if backtracking(i+1, currentSum+(*nums)[i], piecesLeft, expectedSum, used, nums) {
				return true
			}
			(*used)[i] = false
		}
	}
	return false

}

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
	expectedSum := s / k

	used := make([]bool, l)
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	return backtracking(0, 0, k-1, expectedSum, &used, &nums)
}

func CanPartitionKSubsets(nums []int, k int) bool {
	return canPartitionKSubsets(nums, k)
}
