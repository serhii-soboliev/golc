package greedy

/*
659. Split Array into Consecutive Subsequences
https://leetcode.com/problems/split-array-into-consecutive-subsequences/
*/

func isPossible(nums []int) bool {
	n := len(nums)
	shift := 0
	if nums[0] < 0 { shift -= nums[0]}
	
	m := nums[n-1]

	count := make([]int, m+shift+3)
	for i:=0; i<n; i++ {
		nums[i] += shift
		count[nums[i]] += 1
	}

	seq := make([]int, m+shift+1)
	
	for i:=0; i<n; i++ {
		v := nums[i]
		if count[v] == 0 { continue }
		if v == 0 || seq[v-1] == 0 {
			if count[v+1] == 0 || count[v+2] == 0 {
				return false
			} else {
				seq[v+2] +=1
				count[v] -= 1
				count[v+1] -= 1
				count[v+2] -= 1
			}
		} else {
			seq[v-1] -= 1
			seq[v] += 1
			count[v] -= 1
		}
	}
	return true
}

func IsPossible(nums []int) bool {
	return isPossible(nums)
}