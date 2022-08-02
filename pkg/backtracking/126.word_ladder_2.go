package backtracking

/*
126. Word Ladder II
https://leetcode.com/problems/word-ladder-ii/
*/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	found := false
	for _, w := range wordList {
		found = found || (w == endWord)
	}
	res := [][] string {}
	if !found { return res}
	
	notIn := func(seq []string, s string) bool {
		for _, v := range seq {
			if v == s { return false }
		}
		return true
	}
	if notIn(wordList, beginWord) {
		wordList = append(wordList, beginWord)
	}
	minLength := len(wordList) + 1
	dict := make(map[string][]string)

	visited := map[string]bool{beginWord: true}
	for _,v := range wordList {
		visited[v] = false
	}

	for _, s1 := range wordList {
		for _, s2 := range wordList {
			if len(s1) != len(s2) { continue }
			dist := 0
			for i:=0; i<len(s1); i++ {
				if s1[i] != s2[i] { dist++ }
			}
			if dist == 1 {
				dict[s1] = append(dict[s1], s2)	
			}
		}
	}

	var backtrack func(seq []string, current string)

	backtrack = func (seq []string, current string)  {
		if len(seq) > minLength { return }
		if current == endWord { 
			l := len(seq)
			if l == minLength {
				res = append(res, seq)
			}
			if l < minLength {
				minLength = l
				res = [][] string { seq }
			}
		} else {
			for _, v := range dict[current] {
				if !visited[v] {
					cpy := make([]string, len(seq))
					copy(cpy, seq)
					cpy = append(cpy, v)
					visited[v] = true
					backtrack(cpy, v)
					visited[v] = false
				}
			}
		}
	}

	backtrack([]string{beginWord}, beginWord)
	return res
}

func FindLadders(beginWord string, endWord string, wordList []string) [][]string {
	return findLadders(beginWord, endWord, wordList)
}