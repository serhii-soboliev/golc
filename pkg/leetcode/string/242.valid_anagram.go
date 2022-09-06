package string

/*
242. Valid Anagram
https://leetcode.com/problems/valid-anagram/
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := [26]int{}
	for i := 0; i < len(s); i++ {
		chars[int(s[i]-'a')]++
		chars[int(t[i]-'a')]--
	}
	for i := 0; i < 26; i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}

func isAnagramHashMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int)
	for _, v := range s {
		if val, ok := sMap[v]; ok {
			sMap[v] = val + 1
		} else {
			sMap[v] = 1
		}
	}

	tMap := make(map[rune]int)
	for _, v := range t {
		if val, ok := tMap[v]; ok {
			tMap[v] = val + 1
		} else {
			tMap[v] = 1
		}
	}

	for k1, v1 := range sMap {
		if v2, ok := tMap[k1]; ok {
			if v2 != v1 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}

func IsAnagramHashMap(s string, t string) bool {
	return isAnagramHashMap(s, t)
}
