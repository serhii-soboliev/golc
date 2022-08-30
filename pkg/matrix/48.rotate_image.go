package matrix

/*
48. Rotate Image
https://leetcode.com/problems/rotate-image/
*/

func rotate(matrix [][]int)  {
	n := len(matrix)
	for i:=0; i<(n+1)/2; i++ {
		for j:=0; j<n/2; j++ {
			temp := matrix[n - 1 - j][i]
			matrix[n - 1 - j][i] = matrix[n - 1 - i][n - j - 1];
			matrix[n - 1 - i][n - j - 1] = matrix[j][n - 1 -i];
			matrix[j][n - 1 - i] = matrix[i][j];
			matrix[i][j] = temp;
		}
	}
}