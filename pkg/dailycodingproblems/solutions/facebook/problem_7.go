package facebook

/*
This problem was asked by Facebook.

Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.

91. Decode Ways
https://leetcode.com/problems/decode-ways/submissions/
*/

import (
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
    dp := make([]int, len(s) + 1)
    dp[0], dp[1] = 1,1
    for i := 2; i < len(dp); i++ {
        oneD, _ := strconv.Atoi(s[i - 1:i])
		if oneD != 0 { 
			dp[i] += dp[i - 1]
	    }
        twoD, _ := strconv.Atoi(s[i - 2:i])
        if twoD >= 10 && twoD <= 26 { 
			dp[i] += dp[i - 2] 
		}
    }
    return dp[len(s)]
}

func numDecodingsBacktrack(st string) int {

	decode := func(s string) bool {
		if len(s) > 2 || s[0] == '0' {
			return false
		}
		i, err := strconv.Atoi(s)
		if err == nil && i >= 0 && i <= 26 {
			return true	
		} 
		return false
	}

	l := len(st)
	result := 0

	var backtracking func(pos int)

	backtracking = func(pos int) {
		if pos >= l {
			result += 1
			return
		}
		if decode(st[pos:pos+1]) {
			backtracking(pos+1)
		}
		if pos+2<=l && decode(st[pos:pos+2]) {
			backtracking(pos+2)
		}
	}

	backtracking(0)

	return result
}

func NumDecodings(s string) int {
	return numDecodings(s)
}