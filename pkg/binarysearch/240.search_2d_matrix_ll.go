package binarysearch

/*
240. Search a 2D Matrix II
https://leetcode.com/problems/search-a-2d-matrix-ii/
*/
func searchMatrix(matrix [][]int, target int) bool {

	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	
	for i < n && j >= 0 {
		v := matrix[i][j]
		if v == target {
			return true
		} else if v > target {
			j -= 1
		} else { 
			i += 1
		}
		
	}
	return false
}

func SearchMatrix(matrix [][]int, target int) bool {
	return searchMatrix(matrix, target)
}
