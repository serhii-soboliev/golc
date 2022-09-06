package backtracking

/*
526. Beautiful Arrangement
https://leetcode.com/problems/beautiful-arrangement/
*/
func countArrangement(n int) int {
	count := 0
	visited := make([]bool, n+1)

	var backtracking func(int)
	backtracking = func(k int) {
		if k > n {
			count++
		}
		for i := 1; i <= n; i++ {
			if !visited[i] && (i%k == 0 || k%i == 0) {
				visited[i] = true
				backtracking(k + 1)
				visited[i] = false
			}
		}
	}
	backtracking(1)
	return count
}

func CountArrangement(n int) int {
	return countArrangement(n)
}
