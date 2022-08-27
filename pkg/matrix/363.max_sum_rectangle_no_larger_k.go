package matrix

import "math"

/*
363. Max Sum of Rectangle No Larger Than K
https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
*/

func buildPrefixSum(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := [][]int {}
	for i:=0; i<m+1; i++ {
		result = append(result, make([]int, n))
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			result[i+1][j] = result[i][j] + matrix[i][j]
		}
	}
	return result
}

func findRectangleSumClosestToK(prefixSum [][]int, k int) int {
	m, n := len(prefixSum)-1, len(prefixSum[0])
	result := math.MinInt
	for top := 0; top < m; top++ {
		for bottom := top+1; bottom <= m; bottom++ {
			for left:=0; left < n; left++ {
				subresult := 0
				for right := left; right < n; right++ {
					one := prefixSum[bottom][right]
					two := prefixSum[top][right]	
					subresult += one - two	
					if subresult == k {
						return subresult
					}
					if subresult < k && subresult > result {
						result = subresult	
					}
				}
			}
		}
	}
	return result
}


func maxSumSubmatrix(matrix [][]int, k int) int {
	prefixSum := buildPrefixSum(matrix)
	return findRectangleSumClosestToK(prefixSum, k)
}

func MaxSumSubmatrix(matrix [][]int, k int) int {
	return maxSumSubmatrix(matrix, k)
}