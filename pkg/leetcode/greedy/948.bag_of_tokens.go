package greedy

import "sort"

/*
948. Bag of Tokens
https://leetcode.com/problems/bag-of-tokens/
*/

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	n := len(tokens)
	maxScore := 0
	curScore := 0
	lo, hi := 0, n-1
	for lo <= hi && (power >= tokens[lo] || curScore > 0) {
		for lo <= hi && power >= tokens[lo] {
			curScore += 1
			power -= tokens[lo]
			lo += 1
		}
		if curScore > maxScore {
			maxScore = curScore
		}
		if lo <= hi && curScore > 0 {
			power += tokens[hi]
			hi -= 1
			curScore -= 1
		}
	}
	return maxScore
}

func BagOfTokensScore(tokens []int, power int) int {
	return bagOfTokensScore(tokens, power)
}