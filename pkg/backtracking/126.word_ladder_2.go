package backtracking

/*
126. Word Ladder II
https://leetcode.com/problems/word-ladder-ii/
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	
	contains := func(seq []string, s string) bool {
		for _, v := range seq {
			if v == s { return true }
		}
		return false
	}

	areTransformable := func(a, b string) bool {
		diff := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff == 1
	}

	if !contains(wordList, endWord) { return [][]string{}}

	if !contains(wordList, beginWord) {
		wordList = append(wordList, beginWord)
	}

	dict := make(map[string][]string, len(wordList))

	for i, s1 := range wordList {
		for j := i+1; j<len(wordList); j++ {
			s2 := wordList[j]
			if areTransformable(s1, s2) {
				dict[s1] = append(dict[s1], s2)	
				dict[s2] = append(dict[s2], s1)	
			}
		}
	}

	queue, result := [][]string{{beginWord}}, [][]string{}
	pathLen := make(map[string]int, len(wordList))
	pathLen[beginWord] = 1
	for len(queue) > 0 && len(result) == 0 {
		for _, path := range queue {
			queue = queue[1:]
			lastWord := path[len(path) - 1]
			for _, word := range dict[lastWord] {
				if l, ok := pathLen[word]; ok && l < len(path) + 1 {
					continue
				}
				newPath := append(append(make([]string, 0, len(path)+1), path...), word)
				if word == endWord {
					result = append(result, newPath)
				} else {
					queue = append(queue, newPath)
					pathLen[word] = len(newPath)
				}
			}
		}
	}
	return result
}

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	return findLadders(beginWord, endWord, wordList)
}