package string

/*
30. Substring with Concatenation of All Words
https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/

func findSubstring(s string, words []string) []int {

	dict := make(map[string]int)
	for _, word := range words {
		if v, ok := dict[word]; ok {
			dict[word] = v + 1
		} else {
			dict[word] = 1
		}
	}

	n := len(words[0])
	k := len(words)
	l := k * n
	result := []int{}

	check := func(i int) bool {
		wordsCount := 0
		remaining := make(map[string]int)
		for k, v := range dict {
			remaining[k] = v
		}
		for j:=i; j<i+l; j+=n {
			sub := s[j:j+n]
			if v, ok := remaining[sub]; ok && v > 0{
				remaining[sub] = v - 1
				wordsCount += 1	
			} else {
				break
			}	
		}
		return wordsCount == k
	}
	
	for i:=0; i<len(s)-l + 1; i++ {
		if check(i) {
			result = append(result, i)
		}		
	} 
	return result
}

func FindSubstring(s string, words []string) []int {
	return findSubstring(s, words)
}