package dynamicprogramming

import "sort"

/*
823. Binary Trees With Factors
https://leetcode.com/problems/binary-trees-with-factors/
*/

func numFactoredBinaryTrees(a []int) int {
	mod := 1_000_000_007;
	n := len(a)
	dp := make([]int, n)
	index := make(map[int]int)
	sort.Ints(a)
	for i:=0; i<n; i++ {
		dp[i] = 1
		index[a[i]] = i
	}
	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			if a[i] % a[j] == 0 {
				r := a[i] / a[j]
				if v, ok := index[r]; ok {
					dp[i] = (dp[i] + dp[j] * dp[v]) % mod;
				}
			}
		}
	}
	result := 0
	for i:=0; i<n; i++ {
		result += dp[i]
	}
	return result % mod

}

func NumFactoredBinaryTrees(arr []int) int {
	return numFactoredBinaryTrees(arr)
}