package greedy

import "sort"

/*
1647. Minimum Deletions to Make Character Frequencies Unique
https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
*/
const AlphabetLength = 26

func minDeletions(s string) int {
	charCounts := make([]int, AlphabetLength)
	for _, c := range s {
		charCounts[c - 'a'] += 1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(charCounts)))
	delCounts := 0
	for i:=1; i<AlphabetLength; i++ {
		if charCounts[i] == 0 {
			continue
		}
		diff := 0
		if charCounts[i-1] == 0 {
			diff = charCounts[i]
		} else if charCounts[i-1] <= charCounts[i] {
			diff = charCounts[i] - charCounts[i-1] + 1
		}
		charCounts[i] -= diff
		delCounts += diff	
	}
	return delCounts
}

func MinDeletions(s string) int {
	return minDeletions(s)
}