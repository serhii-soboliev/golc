package backtracking

/*
37. Sudoku Solver
https://leetcode.com/problems/sudoku-solver/
*/
func solveSudoku(board [][]byte)  {
	rows,cols,boxes := make([][]bool,9), make([][]bool,9), make([][]bool,9)
    for i := 0; i < 9; i++{
        rows[i], cols[i], boxes[i]= make([]bool,9), make([]bool,9), make([]bool,9)
    }

	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			d := board[i][j]
			if d == '.' { continue }
			num := int(d - '0') - 1		
			boxId := (i / 3) * 3 + j / 3
			rows[i][num], cols[j][num], boxes[boxId][num]  = true, true, true
		}
	}

	solved := false
	
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
		if b == '.' {
			for k:=0;k<=8;k++ {
				if rows[i][k] || cols[j][k] || boxes[boxId][k] {continue}
				uK := byte(k + '1')
				board[i][j] = uK
				rows[i][k], cols[j][k], boxes[boxId][k]  = true, true, true
				backtrack(newI, newJ)
				if !solved {
					board[i][j] = '.'
					rows[i][k], cols[j][k], boxes[boxId][k]  = false, false, false
				} else {
					return 
				}
			} 
		} else {
			backtrack(newI, newJ)
		}
	}

	backtrack(0, 0)
}

func SolveSudoku(board [][]byte)  {
	 solveSudoku(board)
}