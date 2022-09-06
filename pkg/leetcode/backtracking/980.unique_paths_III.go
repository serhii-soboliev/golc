package backtracking

/*
980. Unique Paths III	
https://leetcode.com/problems/unique-paths-iii/
*/

func uniquePathsIII(gr [][]int) int {
	n, m := len(gr), len(gr[0])

	findStartingPoint := func() (int, int) {
		for i:=0; i<n; i++ {
			for j:=0; j<m; j++ {
				if gr[i][j] == 1 {
					return i,j
				}
			}
		}
		panic("No starting point!")
	}

	findPathLength := func () int {
		res := 0
		for i:=0; i<n; i++ {
			for j:=0; j<m; j++ {
				if gr[i][j] == 0 {
					res += 1
				}
			}
		}	
		return res
	}
	
	pathFound := 0
	pathLen := findPathLength() + 1
    dirs := [] struct { x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var backtrack func(int, int, int, [][]int)

	backtrack = func (x int, y int, p int, g [][]int )  {
		if g[x][y] == -1 { return }
		if g[x][y] == 2  { 
			if p == pathLen { pathFound += 1}
			return
		}	
		temp := g[x][y]
		g[x][y] = -1
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < n && 0 <= ny && ny < m {
				backtrack(nx, ny, p+1, g)
			}
		}
		g[x][y] = temp
	}

	centerI, centerJ := findStartingPoint()
	backtrack(centerI, centerJ, 0, gr)
	return pathFound
}

func UniquePathsIII(grid [][]int) int {
	return uniquePathsIII(grid)
}