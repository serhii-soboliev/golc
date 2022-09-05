package backtracking

import "math"

/*
1723. Find Minimum Time to Finish All Jobs
https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/
*/

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	result := math.MaxInt

	workerTimes := make([]int, k)
	
	var backtrack func(pos int, curMax int) 

	backtrack = func(pos int, curMax int) {
		if pos == n {
			result = minInt(result, curMax)
			return
		}
		
		uniqueWorkerTimes := make(map[int]bool)
		for i:=0; i<k; i++ {
			if _, ok := uniqueWorkerTimes[workerTimes[i]]; ok {
				continue;
			}
			uniqueWorkerTimes[workerTimes[i]] = true
			workerTimes[i] += jobs[pos]
			oldCurMax := curMax
			if curMax < workerTimes[i] {
				curMax = workerTimes[i]
			}
			if curMax < result {
				backtrack(pos+1, curMax)
			}
			workerTimes[i] -= jobs[pos]
			curMax = oldCurMax
		}

	}

	backtrack(0, 0)
	return result
}

func minInt(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func MinimumTimeRequired(jobs []int, k int) int {
	return minimumTimeRequired(jobs, k)
}
