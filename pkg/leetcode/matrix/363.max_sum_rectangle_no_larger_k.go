package matrix

import "math"

/*
363. Max Sum of Rectangle No Larger Than K
https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/
*/

func buildPrefixSum(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := [][]int {}
	for i:=0; i<m; i++ {
		result = append(result, make([]int, n+1))
	}
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			result[i][j+1] = result[i][j] + matrix[i][j]
		}
	}
	return result
}

func findRectangleSumClosestToK(prefixSum [][]int, k int) int {
	m, n := len(prefixSum), len(prefixSum[0])
	result := math.MinInt
	for left := 0; left < n-1; left++ {
		for right:=left; right < n-1; right++ {
			for top:=0; top < m; top++ {
				subresult := 0
				for current := top; current < m; current++ {
					subresult += prefixSum[current][right+1] - prefixSum[current][left]
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