package string

/*
890. Find and Replace Pattern
https://leetcode.com/problems/find-and-replace-pattern/
*/

func findAndReplacePattern(words []string, pattern string) []string {
	res := [] string {}

	match := func(s1 string, s2 string) bool {
		m1, m2 := make(map[byte]byte), make(map[byte]byte)
		for i:=0; i<len(s1); i++ {
			b1 := s1[i]
			b2 := s2[i]
			if _, ok := m1[b1]; !ok {
				m1[b1] = b2
			}
			if _, ok := m2[b2]; !ok {
				m2[b2] = b1
			}
			if m1[b1] != b2 || m2[b2] != b1 {
				return false
			}
		}
		return true
	}

	for _, v := range(words) {
		if match(v, pattern) {
			res = append(res, v)
		}
	}

	return res
}

func FindAndReplacePattern(words []string, pattern string) []string {
	return findAndReplacePattern(words, pattern)
}