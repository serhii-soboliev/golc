package string

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