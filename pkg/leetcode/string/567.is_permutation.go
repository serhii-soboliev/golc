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