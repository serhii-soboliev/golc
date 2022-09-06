package backtracking

/*
1593. Split a String Into the Max Number of Unique Substrings	
https://leetcode.com/problems/split-a-string-into-the-max-number-of-unique-substrings/
*/

func maxUniqueSplit(s string) int {
	n := len(s)
	memo := make(map[string]bool)
	result := 0
	
	var backtrack func(pos int, currentResult int) 
	
	backtrack = func(pos int, currentResult int) {
		if pos == n {
			if currentResult > result {
				result = currentResult
			}
			return
		}

		for i:=pos+1; i<=n; i++ {
			newS := s[pos:i]
			if v, ok := memo[newS]; ok && v {
				continue
			}
			memo[newS] = true
			backtrack(i, currentResult + 1)
			memo[newS] = false
		}
	}

	backtrack(0, 0)
	
	return result
}

func MaxUniqueSplit(s string) int {
	return maxUniqueSplit(s)
}