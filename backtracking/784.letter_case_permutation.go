package backtracking

import (
	"strings"
)

/*
784. Letter Case Permutation
https://leetcode.com/problems/letter-case-permutation/
*/

func letterCasePermutation(s string) []string {
	result := []string{}
	word := []byte(strings.ToLower(s))

	var permute func(idx int)

	permute = func(idx int) {
		result = append(result, []string{string(word)}...)
		for i := idx; i < len(word); i++ {
			if 'a' <= word[i] && word[i] <= 'z' {
				word[i] -= 32
				permute(i)
				word[i] += 32
			}
		}
	}
	permute(0)
	return result
}

func LetterCasePermutation(s string) []string {
	return letterCasePermutation(s)
}

func isUpperCaseLetter(symbol byte) bool {
	return symbol >= 'A' && symbol <= 'Z'
}

func isDigit(symbol byte) bool {
	return '0' <= symbol && symbol <= '9'
}

func permute(word []byte, idx int, ans *[]string) {
	if idx == len(word) {
		*ans = append(*ans, string(word))
		return
	}
	currentSymbol := word[idx]
	newIdx := idx + 1
	if isDigit(currentSymbol) {
		permute(word, newIdx, ans)
		return
	}
	var newSymbol byte
	if isUpperCaseLetter(currentSymbol) {
		newSymbol = currentSymbol + 32
	} else {
		newSymbol = currentSymbol - 32
	}
	var oldWord = make([]byte, len(word))
	copy(oldWord, word)
	var newWord = append(word[:idx], append([]byte{newSymbol}, word[newIdx:]...)...)
	permute(oldWord, newIdx, ans)
	permute(newWord, newIdx, ans)
}

func letterCasePermutationDFS(S string) []string {
	ans, word := []string{}, []byte(S)
	permute(word, 0, &ans)
	return ans
}

func LetterCasePermutationDFS(s string) []string {
	return letterCasePermutationDFS(s)
}
