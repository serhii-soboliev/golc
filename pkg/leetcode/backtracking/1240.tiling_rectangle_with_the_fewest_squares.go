package backtracking

/*
1240. Tiling a Rectangle with the Fewest Squares
https://leetcode.com/problems/tiling-a-rectangle-with-the-fewest-squares/
*/

func tilingRectangleBacktrack(n int, m int) int {

	if m == n {
		return 1
	}

	min := func(a []int) (int, int) {
		var m, idx int
		for i := 0; i < len(a); i++ {
			if i == 0 || a[i] < m {
				m = a[i]
				idx = i
			}
		}
		return m, idx
	}

	minSquareCount := n * m

	h := make([]int, m)

	var backtrack func(int)

	backtrack = func(squareCount int) {
		if squareCount >= minSquareCount {
			return
		}

		mH, start := min(h)
		if mH == n {
			if minSquareCount > squareCount {
				minSquareCount = squareCount
			}
			return
		}
		end := start
		for end < m && h[end] == h[start] && end-start+1 <= n-mH {
			end += 1
		}
		for i := end - 1; i >= start; i-- {
			side := i - start + 1
			for j := start; j <= i; j++ {
				h[j] += side
			}
			backtrack(squareCount + 1)
			for j := start; j <= i; j++ {
				h[j] -= side
			}
		}

	}

	backtrack(0)
	return minSquareCount
}

func TilingRectangleBacktrack(n int, m int) int {
	return tilingRectangleBacktrack(n, m)
}
