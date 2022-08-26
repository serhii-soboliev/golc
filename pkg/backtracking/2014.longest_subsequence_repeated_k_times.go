package backtracking

import "sort"

/*
2014. Longest Subsequence Repeated k Times
https://leetcode.com/problems/longest-subsequence-repeated-k-times/
*/

func KCombinations(a []byte, k int) (result [][]byte) {
	
	var combBacktrack func(pos int,l int, current []byte)

	combBacktrack = func(pos int, l int, current []byte) {
		if l == 0 {
			result = append(result, current)
			return
		}
		currendIdx := k - l
		for i:=pos; i<=len(a)-l; i++ {
			cpy := make([]byte, len(current))
			copy(cpy, current)
			cpy[currendIdx] = a[i]
			combBacktrack(i+1, l-1, cpy)
		}
	}

	combBacktrack(0, k, make([]byte, k))

	return
}

func Perm(a []byte) (result []string) {

	resMap := make(map[string]bool)

	var perm func(a []byte,  i int) 

	perm = func(a []byte, i int) {
		if i > len(a) {
			resMap[string(a)] = true
			return 
		}
		perm(a, i+1)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			perm(a, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}

	perm(a, 0)
	for k := range resMap {
		result = append(result, k)
	}
	return
}

func doesMeetKTimes(a, b string, times int) bool {
	l := len(a)
	totalTimes := times * l
	currentTimes := 0
	for i:=0; i<len(b); i++ {
		if b[i] == a[currentTimes % l] {
			currentTimes += 1
			if currentTimes == totalTimes { 
				return true
			}
		}
	}
	return false
}

func PossibleKLenSubStrings(a []byte, k int) (result []string) {
	combinations := KCombinations(a, k)
	for _, combination := range combinations {
		result = append(result, Perm(combination)...)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(result)))
	return
} 

func longestSubsequenceRepeatedK(s string, k int) string {
	chars := make([]int, 26)
	for _, c := range s {
		chars[c - 'a'] += 1
	}

	possibleCharsForKLenStrings := func(n int) (possibleChars []byte) {
		for i, c := range chars {
			if c >= n {
				times := c / n
				for j:=0; j<times; j++ {
					possibleChars = append(possibleChars, byte(i + 'a'))
				}	
			}
		}

		return
	}

	possibleChars := possibleCharsForKLenStrings(k)

	for i:=len(possibleChars); i>=1; i-- {
		found := false
		possibleSubStrings := PossibleKLenSubStrings(possibleChars, i)

		for _, substr := range possibleSubStrings {
			if doesMeetKTimes(substr, s, k) {
				return substr
			}
		}
		if found {
			break
		}
	}	

	return "" 
}

func LongestSubsequenceRepeatedK(s string, k int) string {
	return longestSubsequenceRepeatedK(s, k)
}