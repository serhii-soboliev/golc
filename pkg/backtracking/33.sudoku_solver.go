package backtracking

/*
37. Sudoku Solver
https://leetcode.com/problems/sudoku-solver/
*/
func solveSudoku(board [][]byte)  {
    rows := [9] map[byte]bool {}
	cols := [9] map[byte]bool {}
	boxes := [9] map[byte]bool {}
	solved := false

	for i:=0; i<9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			d := board[i][j]
			if d == '.' { continue }
			rows[i][d] = true
			cols[j][d] = true
			boxId := (i / 3) * 3 + j / 3
			boxes[boxId][d] = true
		}
	}

	isIn := func(m map[byte]bool, b byte) bool {
		if val, ok := m[b]; ok {
			return val
		} else {
			return false
		}
	}
	
	var backtrack func(i int, j int)
	backtrack = func (i int, j int)  {
		if i == 9 {
			solved = true
			return
		}
		newI := i + (j + 1) / 9
		newJ := (j + 1) % 9
		boxId := (i / 3) * 3 + j / 3
		b := board[i][j]
		if b != '.' {
			backtrack(newI, newJ)
		} else {
			for k:=1;k<=9;k++ {
				uK := byte(k + '0')
				if isIn(rows[i], uK) || isIn(cols[j], uK) || isIn(boxes[boxId], uK) {continue}
				board[i][j] = uK
				rows[i][uK] = true
				cols[j][uK] = true
				boxes[boxId][uK] = true
				backtrack(newI, newJ)
				if !solved {
					board[i][j] = '.'
					rows[i][uK] = false
					cols[j][uK] = false
					boxes[boxId][uK] = false	
				} else {
					return 
				}
			} 
		}
	}

	backtrack(0, 0)
}

func SolveSudoku(board [][]byte)  {
	 solveSudoku(board)
}