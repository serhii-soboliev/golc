package string

/*
792. Number of Matching Subsequences
https://leetcode.com/problems/number-of-matching-subsequences/
*/
func numMatchingSubseq(s string, words []string) int {
	wordDict := map[byte][]string{}
	for _, w := range words {
		wordDict[w[0]] = []string{}
	}
	for _, w := range words {
		wordDict[w[0]] = append(wordDict[w[0]], []string{w}...)
	}
	count := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		expected := wordDict[c]
		wordDict[c] = []string{}
		for _, w := range expected {
			if len(w) == 1 {
				count += 1
			} else {
				newValue := append(wordDict[w[1]], []string{w[1:]}...)
				wordDict[w[1]] = newValue
			}
		}
	}
	return count
}

func NumMatchingSubseq(s string, words []string) int {
	return numMatchingSubseq(s, words)
}
