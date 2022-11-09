package string

/*
567. Permutation in String
https://leetcode.com/problems/permutation-in-string/description/
*/

func checkInclusion(s1 string, s2 string) bool {
	const AlphabetLength = 26
	s1Counts := make([]int, AlphabetLength)
	for _, c := range s1 {
		s1Counts[c - 'a'] += 1
	}
	count := len(s1)
	s2Chars := []rune(s2)
	left, right := 0,0
	for right < len(s2Chars) {
		rIdx := s2Chars[right] - 'a'
		if s1Counts[rIdx] > 0 {	
			count -= 1	
		}
		s1Counts[rIdx] -= 1
		right += 1	
		for count == 0 {
			if right - left == len(s1) {
				return true
			}
			s1Counts[s2Chars[left] - 'a'] += 1
			if s1Counts[s2Chars[left] - 'a'] > 0 {
				count += 1
			} 
			left += 1
		}
	}
	return false
	
}

func CheckInclusion(s1 string, s2 string) bool {
	return checkInclusion(s1, s2)
}

/*
1.2
Check Permutation: Given two strings,write a method to decide if one is a permutation of the
other.
*/

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	const AlphabetLength = 26
	s1Counts := make([]int, AlphabetLength)
	for _, c := range s1 {
		s1Counts[c - 'a'] += 1
	}
	count := len(s1) 
	for _, c := range s2 {
		idx := c - 'a'
		if s1Counts[idx] == 0 {
			return false
		}
		s1Counts[idx] -= 1
		count -= 1
	}
	return true
}