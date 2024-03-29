package string

/*
5. Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/
*/
func LongestCommonSubstring(s1 string, s2 string) []string {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, l1)
	result := []string{}
	maxLen := 0
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if s1[i] == s2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					result = []string{s1[i-maxLen+1 : i+1]}
				} else if dp[i][j] == maxLen {
					result = append(result, s1[i-maxLen+1:i+1])
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return result
}

func longestPalindromeEAC(s string) string {
	n := len(s)

	expandAroundCenter := func(i, j int) int {
		for 0 <= i && j < n && s[i] == s[j] {
			i--
			j++
		}
		return j - i - 1
	}

	start, end := 0, 0

	for i := 0; i < n; i++ {
		l1 := expandAroundCenter(i, i)
		if l1 > end-start+1 {
			start = i - l1/2
			end = i + l1/2
		}
		l2 := expandAroundCenter(i, i+1)
		if l2 > end-start+1 {
			start = i - (l2-1)/2
			end = i + 1 + (l2-1)/2
		}
	}
	return s[start : end+1]
}

func longestPalindromeDP(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i:=0; i<n; i++ {
		dp[i] = make([]int, n)
	}
	for i:=0; i<n; i++ {
		dp[i][i] = 1
	}
	for i:=0; i<n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 0
		}
	}
	for i:=2; i<=n; i++ {
		for j:=0; j<n-i; j++ {
			if dp[j+1][j+i-1] > 0 && s[j] == s[j+i]{
				dp[j][j+i] = dp[j+1][j+i-1] + 2
			}	
		}
	}
	maxLen := 0
	result := ""
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
				result = s[i:j+1]
			}
		}

	}
	return result
}

func LongestPalindromeEAC(s string) string {
	return longestPalindromeEAC(s)
}

func LongestPalindromeDP(s string) string {
	return longestPalindromeDP(s)
}
