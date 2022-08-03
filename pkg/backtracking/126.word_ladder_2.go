package backtracking

import "fmt"

/*
126. Word Ladder II
https://leetcode.com/problems/word-ladder-ii/
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	buildReverseDict := func() map[string][]string {
		adjacencyList := make(map[string][]string)
		for _, word := range wordList {
			for i := 0; i < len(word); i++ {
				pattern := fmt.Sprintf("%s.%s", word[:i], word[i+1:])
				adjacencyList[pattern] = append(adjacencyList[pattern], word)
			}
		}

		reverseDict := map[string][]string{beginWord: {}}

		q := []string{beginWord}

		for len(q) > 0  {
			n := len(q)
			visitedThisLevel := make(map[string][]string)
			for i := 0; i < n; i++ {
				word := q[0]
				q = q[1:]
				for i := 0; i < len(word); i++ {
					pattern := fmt.Sprintf("%s.%s", word[:i], word[i+1:])
					for _, nextWord := range adjacencyList[pattern] {
						if _, ok1 := reverseDict[nextWord]; !ok1 {
							if _, ok2 := visitedThisLevel[nextWord]; !ok2 {
								visitedThisLevel[nextWord] = []string{word}
								q = append(q, nextWord)
							} else {
								visitedThisLevel[nextWord] = append(visitedThisLevel[nextWord], word)
							}
						}
					}
				}

			}

			for k, v := range visitedThisLevel {
				reverseDict[k] = append(reverseDict[k], v...)
			}

		}
		return reverseDict
	}

	contains := func(seq []string, s string) bool {
		for _, v := range seq {
			if v == s {
				return true
			}
		}
		return false
	}

	if !contains(wordList, endWord) {
		return [][]string{}
	}

	reverseDict := buildReverseDict()

	buildWordLadders := func() [][]string {
		var dfs func(node string) [][]string
		dfs = func(node string) [][]string {
			if node == beginWord {
				return [][]string{{beginWord}}
			}
			if _, ok := reverseDict[node]; !ok {
				return [][]string{}
			}
			result := [][]string{}

			for _, parent := range reverseDict[node] {
				parentPathes := dfs(parent)
				for i := 0; i < len(parentPathes); i++ {
					result = append(result, parentPathes[i])
				}
				if node == endWord {
					fmt.Print(endWord)
				}
				for i := len(result) - 1; i > len(result)-len(parentPathes)-1; i-- {
					result[i] = append(result[i], node)
				}
			}

			return result
		}
		return dfs(endWord)
	}

	return buildWordLadders()
}

func findLaddersBFS(beginWord string, endWord string, wordList []string) [][]string {

	contains := func(seq []string, s string) bool {
		for _, v := range seq {
			if v == s {
				return true
			}
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

	if !contains(wordList, endWord) {
		return [][]string{}
	}

	if !contains(wordList, beginWord) {
		wordList = append(wordList, beginWord)
	}

	dict := make(map[string][]string, len(wordList))

	for i, s1 := range wordList {
		for j := i + 1; j < len(wordList); j++ {
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
			lastWord := path[len(path)-1]
			for _, word := range dict[lastWord] {
				if l, ok := pathLen[word]; ok && l < len(path)+1 {
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

func FindLaddersBFS(beginWord string, endWord string, wordList []string) [][]string {
	return findLaddersBFS(beginWord, endWord, wordList)
}
