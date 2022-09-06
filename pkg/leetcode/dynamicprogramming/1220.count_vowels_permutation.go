package dynamicprogramming

/*
1220. Count Vowels Permutation
https://leetcode.com/problems/count-vowels-permutation/
*/

func countVowelPermutationS(n int) int {
	const mod = 1_000_000_007
	a, e, i, o, u := 1, 1, 1, 1, 1
	for j := 0; j < n-1; j++ {
		a, e, i, o, u = (e+i+u)%mod, (a+i)%mod, (e+o)%mod, i, (o+i)%mod
	}
	return (a + e + i + o + u) % mod
}

func countVowelPermutationD(n int) int {
	const mod = 1_000_000_007
	dp := [5][]int{}
	for i:=0; i<5; i++ {
		dp[i] = make([]int, n)
	}
	for i:=0; i<5; i++ {
		dp[i][0] = 1
	}
	for i:=1;i<n;i++ {
		dp[0][i] = (dp[1][i-1] + dp[2][i-1] + dp[4][i-1]) % mod
		dp[1][i] = (dp[0][i-1] + dp[2][i-1]) % mod
		dp[2][i] = (dp[1][i-1] + dp[3][i-1]) % mod
		dp[3][i] = dp[2][i-1] % mod 
		dp[4][i] = (dp[2][i-1] + dp[3][i-1]) % mod
	}
	result := 0
	for i:=0; i<5; i++ {
		result += dp[i][n-1]
	}
	return result
}

func CountVowelPermutationS(n int) int {
	return countVowelPermutationS(n)
}

func CountVowelPermutationD(n int) int {
	return countVowelPermutationD(n)
}