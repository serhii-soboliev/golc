package backtracking

import "sort"

/*
1723. Find Minimum Time to Finish All Jobs
https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/
*/

const MaxInt = int(^uint(0) >> 1)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	result := MaxInt
	sort.Sort(sort.Reverse(sort.IntSlice(jobs)))
	workerTimes := make([]int, k)
	
	var backtrack func(pos int, curMax int) 

	backtrack = func(pos int, curMax int) {
		if pos == n {
			if result > curMax {
				result = curMax
			}
			return
		}

		jobDuration := jobs[pos]
		uniqueWorkerTimes := make(map[int]bool)
		for i:=0; i<k; i++ {
			if _, ok := uniqueWorkerTimes[workerTimes[i]]; ok {
				continue;
			}
			uniqueWorkerTimes[workerTimes[i]] = true
			workerTimes[i] += jobDuration
			oldCurMax := curMax
			if curMax < workerTimes[i] {
				curMax = workerTimes[i]
			}
			if curMax < result {
				backtrack(pos+1, curMax)
			}
			workerTimes[i] -= jobDuration
			curMax = oldCurMax
			if workerTimes[i] == 0 {
				break
			}
		}

	}

	backtrack(0, 0)
	return result
}


func MinimumTimeRequired(jobs []int, k int) int {
	return minimumTimeRequired(jobs, k)
}
