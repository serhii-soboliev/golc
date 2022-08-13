package string

/*
30. Substring with Concatenation of All Words
https://leetcode.com/problems/substring-with-concatenation-of-all-words/
*/

func findSubstringMap(s string, words []string) []int {

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
		for j := i; j < i+l; j += n {
			sub := s[j : j+n]
			if v, ok := remaining[sub]; ok && v > 0 {
				remaining[sub] = v - 1
				wordsCount += 1
			} else {
				break
			}
		}
		return wordsCount == k
	}

	for i := 0; i < len(s)-l+1; i++ {
		if check(i) {
			result = append(result, i)
		}
	}
	return result
}

func findSubstringSlidingWindow(s string, words []string) []int {
	dict := make(map[string]int)
	for _, word := range words {
		dict[word] += 1
	}

	wordLen, wordsCount := len(words[0]), len(words)
	substringSize := wordsCount * wordLen
	result := []int{}

	slidingWindow := func(left int) {
		wordsFound, excessWord, wordsUsed := make(map[string]int), false, 0

		for right := left; right + wordLen <= len(s); right += wordLen {
			sub := s[right:right+wordLen]
			if v, ok1 := dict[sub]; !ok1 {
				wordsFound, excessWord, wordsUsed, left = make(map[string]int), false, 0, right + wordLen				
			} else {
				for right - left == substringSize || excessWord {
					leftmostWord := s[left:left+wordLen]
					left += wordLen
					wordsFound[leftmostWord] -= 1
					if wordsFound[leftmostWord] >= dict[leftmostWord] {
						excessWord = false
					} else {
						wordsUsed -= 1
					}
				}

				wordsFound[sub] += 1

				if wordsFound[sub] <= v {
					wordsUsed += 1
				} else {
					excessWord = true
				}
				if wordsUsed == wordsCount && !excessWord {
					result = append(result, left)
				}
			}
		}	
	}
	for i:=0; i<wordLen; i++ {
		slidingWindow(i)
	}
	return result
}

func FindSubstringSlidingWindow(s string, words []string) []int {
	return findSubstringSlidingWindow(s, words)
}

func FindSubstringMap(s string, words []string) []int {
	return findSubstringMap(s, words)
}
