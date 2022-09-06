package matrix

/*
1329. Sort the Matrix Diagonally
https://leetcode.com/problems/sort-the-matrix-diagonally/
*/

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])

	inBound := func(i int, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	bucketSort := func(a []int) {
		s := [101]int{}
		for _, v := range a {
			s[v] += 1
		}
		k := 0
		for i, v := range s {
			for j:=0; j<v; j++ {
				a[k] = i
				k += 1
			}
		}

	}

	sortDiagonal := func (i int, j int) {
		new_i, new_j := i, j
		diagonal := []int{}
		for inBound(new_i, new_j) {
			diagonal = append(diagonal, mat[new_i][new_j])
			new_i += 1
			new_j += 1
		}
		bucketSort(diagonal)
		new_i, new_j = i, j
		for _, v := range diagonal {
			mat[new_i][new_j] = v
			new_i += 1
			new_j += 1
		}
	}

	for i:=0; i<m-1; i++ {
		sortDiagonal(i,0)
	}

	for i:=1; i<n-1; i++ {
		sortDiagonal(0,i)
	}

	return mat
}

func DiagnoalSort(mat [][]int) [][]int {
	return diagonalSort(mat)
}