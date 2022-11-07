package string

import "sort"

/*
387. First Unique Character in a String
https://leetcode.com/problems/first-unique-character-in-a-string/
*/

func firstUniqChar(s string) int {
	n := len(s)
	dict := make(map[rune]int)
	freq := make([]int, n)
	for i, c := range s {
		if v, ok := dict[c]; ok {
			freq[v] = 0
		} else {
			freq[i] = 1
			dict[c] = i
		}	
	}
	for i, v := range freq {
		if v == 1 {
			return i
		}
	}
	return -1
}

/*
1.1
Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/
func IsUnique(word string) bool {
	n := len(word)
	s := []rune(word)
    sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })

	t := s[0]
	for i:=1; i<n; i++ {
		if s[i] == t {
			return false
		}
		t = s[i]	
	}
	return true
}

func IsUnique2(word string) bool {
	checker := 0
	for _, v := range word {
		val := int(v - 'a')
		b := 1 << val
		if checker & b > 0 {
			return false
		}
		checker |= b	
	}
	return true
}


